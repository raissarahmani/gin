package repositories

import (
	"context"
	"main/internal/models"

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

func (s *ShowingRepositories) BookSchedule(ctx context.Context, movieID, cityID, cinemaID, scheduleID int) ([]models.Schedule, error) {
	query := `
		SELECT sc.book_date AS date, sc.book_time AS time, ct.city, c.cinema_name, m.title
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
		if err := rows.Scan(&s.Date, &s.Time, &s.City, &s.Cinema, &s.Movie); err != nil {
			return nil, err
		}
		schedules = append(schedules, s)
	}

	return schedules, nil
}

func (s *ShowingRepositories) GetSchedulesByMovie(ctx context.Context, movieID int) ([]models.Schedule, error) {
	query := `
		SELECT DISTINCT sc.id, sc.book_date AS date, sc.book_time AS time
		FROM showing_schedule sh
		JOIN schedule sc ON sh.schedule_id = sc.id
		WHERE sh.movies_id = $1
		ORDER BY sc.book_date, sc.book_time`
	rows, err := s.db.Query(ctx, query, movieID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var schedules []models.Schedule
	for rows.Next() {
		var s models.Schedule
		if err := rows.Scan(&s.Id, &s.Date, &s.Time); err != nil {
			return nil, err
		}
		schedules = append(schedules, s)
	}

	return schedules, nil
}

func (s *ShowingRepositories) GetCitiesByMovie(ctx context.Context, movieID int) ([]models.Schedule, error) {
	query := `
		SELECT DISTINCT ct.id, ct.city
		FROM showing_schedule sh
		JOIN city ct ON sh.city_id = ct.id
		WHERE sh.movies_id = $1
		ORDER BY ct.id`
	rows, err := s.db.Query(ctx, query, movieID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cities []models.Schedule
	for rows.Next() {
		var s models.Schedule
		if err := rows.Scan(&s.Id, &s.City); err != nil {
			return nil, err
		}
		cities = append(cities, s)
	}

	return cities, nil
}

func (s *ShowingRepositories) GetCinemasByFilters(ctx context.Context, movieID, cityID, scheduleID int) ([]models.Schedule, error) {
	query := `
		SELECT DISTINCT c.id, c.cinema_name, c.price, sh.schedule_id
		FROM showing_schedule sh
		JOIN cinema c ON sh.cinema_id = c.id
		WHERE sh.movies_id = $1 AND sh.city_id = $2 AND sh.schedule_id = $3
		ORDER BY c.cinema_name`
	rows, err := s.db.Query(ctx, query, movieID, cityID, scheduleID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cinemas []models.Schedule
	for rows.Next() {
		var s models.Schedule
		if err := rows.Scan(&s.Id, &s.Cinema, &s.Price, &s.ScheduleID); err != nil {
			return nil, err
		}
		cinemas = append(cinemas, s)
	}
	return cinemas, nil
}

func (s *ShowingRepositories) GetSeatAvailability(ctx context.Context, movieID, cityID, cinemaID, scheduleID int) ([]models.Seat, error) {
	query := `
		SELECT s.id AS seat_id, s.seat_number AS seat, ss.is_available
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
		WHERE ss.movies_id = $1 AND ss.city_id = $2 AND ss.cinema_id = $3 AND ss.schedule_id = $4
		ORDER BY 
		  LEFT(s.seat_number, 1),
		  CAST(SUBSTRING(s.seat_number FROM 2) AS INTEGER)`

	rows, err := s.db.Query(ctx, query, movieID, cityID, cinemaID, scheduleID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var seats []models.Seat
	for rows.Next() {
		var sch models.Seat
		err := rows.Scan(&sch.Id, &sch.Seat_number, &sch.Is_available)
		if err != nil {
			return nil, err
		}
		seats = append(seats, sch)
	}

	return seats, nil
}
