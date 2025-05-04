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

func (s *ShowingRepository) GetSchedulesByMovieID(ctx context.Context, movieID, cityID, cinemaID int) ([]models.Schedule, error) {
	query := `
		SELECT 
			sc.book_date AS date,
			sc.book_time AS time,
			ct.city,
			c.cinema_name,
			m.title
		FROM showing_schedule sh
		JOIN schedule sc ON sh.schedule_id = sc.id
		JOIN movies m ON sh.movies_id = m.id
		JOIN cinema c ON sh.cinema_id = c.id
		JOIN city ct ON sh.city_id = ct.id
		WHERE sh.movies_id = $1 AND sh.city_id = $2 AND sh.cinema_id = $3
	`

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
