package repositories

import (
	"context"
	"errors"
	"fmt"
	"main/internal/models"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ShowingRepositories struct {
	db *pgxpool.Pool
}

func NewShowingRepository(pg *pgxpool.Pool) *ShowingRepositories {
	return &ShowingRepositories{
		db: pg,
	}
}

func (s *ShowingRepositories) GetSchedulesByMovieID(ctx context.Context, movieID, cityID, cinemaID int) ([]models.Schedule, error) {
	query := `
		SELECT sh.id, sc.book_date AS date, sc.book_time AS time, ct.city, c.cinema_name, m.title
		FROM showing_schedule sh
		JOIN schedule sc ON sh.schedule_id = sc.id
		JOIN movies m ON sh.movies_id = m.id
		JOIN cinema c ON sh.cinema_id = c.id
		JOIN city ct ON sh.city_id = ct.id
		WHERE sh.movies_id = $1 AND sh.city_id = $2 AND sh.cinema_id = $3`

	rows, err := s.db.Query(ctx, query, movieID, cityID, cinemaID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var schedules []models.Schedule
	for rows.Next() {
		var s models.Schedule
		if err := rows.Scan(&s.Id, &s.Date, &s.Time, &s.City, &s.Cinema, &s.Movie); err != nil {
			return nil, err
		}
		schedules = append(schedules, s)
	}

	return schedules, nil
}

func (s *ShowingRepositories) GetSeatAvailability(ctx context.Context, movieID, cityID, cinemaID, scheduleID int) ([]models.Schedule, error) {
	query := `
	SELECT sc.book_date AS date, sc.book_time AS time, ct.city, c.cinema_name, m.title, s.seat_number AS seat, ss.is_available
	FROM showing_seat_schedule ss
	JOIN seat s ON ss.seat_id = s.id
	JOIN showing_schedule sh ON 
		ss.schedule_id = sh.schedule_id AND 
		ss.movies_id = sh.movies_id AND 
		ss.cinema_id = sh.cinema_id AND 
		ss.city_id = sh.city_id
	JOIN schedule sc ON sh.schedule_id = sc.id
	JOIN movies m ON sh.movies_id = m.id
	JOIN cinema c ON sh.cinema_id = c.id
	JOIN city ct ON sh.city_id = ct.id
	WHERE ss.movies_id = $1 AND ss.city_id = $2 AND ss.cinema_id = $3 AND ss.schedule_id = $4 AND ss.is_available = false
	ORDER BY s.seat_number`

	rows, err := s.db.Query(ctx, query, movieID, cityID, cinemaID, scheduleID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var seats []models.Schedule
	for rows.Next() {
		var sch models.Schedule
		err := rows.Scan(&sch.Seat)
		if err != nil {
			return nil, err
		}
		seats = append(seats, sch)
	}

	return seats, nil
}

func (s *ShowingRepositories) IsSeatAvailable(ctx context.Context, movieID, cityID, cinemaID, scheduleID, seatID int) (bool, error) {
	query := `
	SELECT 1
	FROM showing_seat_schedule ss
	WHERE ss.movies_id = $1 AND ss.city_id = $2 AND ss.cinema_id = $3 AND ss.schedule_id = $4 AND ss.seat_id = $5 AND ss.is_available = false
	LIMIT 1`

	var seat int
	err := s.db.QueryRow(ctx, query, movieID, cityID, cinemaID, scheduleID, seatID).Scan(&seat)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return true, nil
		}
		return false, err
	}

	return false, nil
}

func (s *ShowingRepositories) BookSeat(ctx context.Context, seatID, movieID, cityID, cinemaID, scheduleID int) error {
	query := `
		UPDATE showing_seat_schedule 
		SET is_available = false 
		WHERE seat_id = $1 AND movies_id = $2 AND city_id = $3 AND cinema_id = $4 AND schedule_id = $5 AND is_available = true`

	res, err := s.db.Exec(ctx, query, seatID, movieID, cityID, cinemaID, scheduleID)
	if err != nil {
		return err
	}

	if res.RowsAffected() == 0 {
		return fmt.Errorf("seat is already taken")
	}

	return nil
}
