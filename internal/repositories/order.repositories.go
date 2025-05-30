package repositories

import (
	"context"
	"fmt"
	"log"
	"main/internal/models"
	"strings"

	"github.com/jackc/pgx/v5/pgxpool"
)

type OrderRepositories struct {
	db *pgxpool.Pool
}

func NewOrderRepository(pg *pgxpool.Pool) *OrderRepositories {
	return &OrderRepositories{
		db: pg,
	}
}

func (o *OrderRepositories) IsSeatAvailable(ctx context.Context, movieID, cityID, cinemaID, scheduleID int, seatIDs []int) (bool, error) {
	query := `
	SELECT COUNT (*)
	FROM showing_seat_schedule ss
	WHERE ss.movies_id = $1 AND ss.city_id = $2 AND ss.cinema_id = $3 AND ss.schedule_id = $4 AND ss.seat_id = ANY($5) AND ss.is_available = false`

	var seats int
	err := o.db.QueryRow(ctx, query, movieID, cityID, cinemaID, scheduleID, seatIDs).Scan(&seats)
	if err != nil {
		return false, err
	}
	return seats == 0, nil
}

func (o *OrderRepositories) GetTicketPrice(ctx context.Context, cinemaID int) (int, error) {
	var price int
	err := o.db.QueryRow(ctx, `SELECT price FROM cinema WHERE id = $1`, cinemaID).Scan(&price)
	return price, err
}

func (o *OrderRepositories) BookOrder(ctx context.Context, userID int, input models.OrderRequest, totalPrice int) error {
	tx, err := o.db.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	// Check seats availability
	log.Printf("[DEBUG] Booking request for seats: %v", input.SeatIDs)
	rows, err := tx.Query(ctx, `
		SELECT seat_id FROM showing_seat_schedule
		WHERE schedule_id = $1 AND city_id = $2 AND cinema_id = $3 AND movies_id = $4
		AND seat_id = ANY($5) AND is_available = true
		FOR UPDATE`,
		input.ScheduleID, input.CityID, input.CinemaID, input.MovieID, input.SeatIDs)
	if err != nil {
		return err
	}
	defer rows.Close()

	availableSeatIDs := make(map[int]bool)
	for rows.Next() {
		var seatID int
		if err := rows.Scan(&seatID); err != nil {
			return err
		}
		log.Printf("[DEBUG] Seat %d is available", seatID)
		availableSeatIDs[seatID] = true
	}

	if len(availableSeatIDs) != len(input.SeatIDs) {
		log.Printf("[DEBUG] Mismatch in available seats: got %d available, expected %d", len(availableSeatIDs), len(input.SeatIDs))

		missingSeats := []int{}
		for _, sid := range input.SeatIDs {
			if !availableSeatIDs[sid] {
				missingSeats = append(missingSeats, sid)
			}
		}
		log.Printf("[DEBUG] These seats are not available: %v", missingSeats)

		return fmt.Errorf("some seats are already booked")
	}

	// Insert transaction
	var transactionID int
	err = tx.QueryRow(ctx, `
		INSERT INTO transaction (users_id, fullname, email, phone, total_price, payment_method_id, payment_done)
		VALUES ($1, $2, $3, $4, $5, $6, true)
		RETURNING id`, userID, input.Fullname, input.Email, input.Phone, totalPrice, input.PaymentMethodID).Scan(&transactionID)
	if err != nil {
		return err
	}

	// Insert into seat_selection
	query := `INSERT INTO seat_selection (transaction_id, schedule_id, city_id, cinema_id, movies_id, seat_id) VALUES `
	values := []any{}
	placeholders := []string{}

	for i, seatID := range input.SeatIDs {
		start := i*6 + 1
		placeholders = append(placeholders,
			fmt.Sprintf("($%d,$%d,$%d,$%d,$%d,$%d)", start, start+1, start+2, start+3, start+4, start+5),
		)
		values = append(values,
			transactionID,
			input.ScheduleID,
			input.CityID,
			input.CinemaID,
			input.MovieID,
			seatID,
		)
	}

	query += strings.Join(placeholders, ",")

	log.Println("[DEBUG]: query", query)
	_, err = tx.Exec(ctx, query, values...)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	// Update seat availability
	_, err = tx.Exec(ctx, `
		UPDATE showing_seat_schedule
		SET is_available = false
		WHERE schedule_id = $1 AND city_id = $2 AND cinema_id = $3 AND movies_id = $4
		AND seat_id = ANY($5)`, input.ScheduleID, input.CityID, input.CinemaID, input.MovieID, input.SeatIDs)
	if err != nil {
		return err
	}

	return tx.Commit(ctx)
}

func (o *OrderRepositories) GetHistory(ctx context.Context, userID int) ([]models.OrderHistory, error) {
	query := `
	SELECT
	  t.id AS transaction_id, 
	  sch.book_date, sch.book_time, ct.city, c.cinema_name, m.title, s.seat_number
	FROM transaction t
	JOIN seat_selection ss ON ss.transaction_id = t.id
	JOIN seat s ON s.id = ss.seat_id
	JOIN schedule sch ON sch.id = ss.schedule_id
	JOIN movies m ON m.id = ss.movies_id
	JOIN cinema c ON c.id = ss.cinema_id
	JOIN city ct ON ct.id = ss.city_id
	WHERE t.users_id = $1
	ORDER BY t.id, s.seat_number`

	rows, err := o.db.Query(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []models.OrderHistory
	for rows.Next() {
		var order models.OrderHistory
		err := rows.Scan(
			&order.TransactionID, &order.BookDate, &order.BookTime, &order.City, &order.CinemaName, &order.MovieTitle, &order.SeatNumber)
		if err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}

	return orders, nil
}
