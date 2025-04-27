package repositories

import (
	"context"
	"main/internal/models"
	"main/pkg"
)

type ScheduleRepositories struct{}

var ScheduleRepo *ScheduleRepositories

func NewScheduleRepository() {
	ScheduleRepo = &ScheduleRepositories{}
}

func (s *ScheduleRepositories) GetSchedulesByMovieID(c context.Context, movie_id int) ([]models.Schedule, error) {
	query := `
		SELECT o.id, m.title, o.book_date, ot.book_time, ol.location 
		FROM "order" o
		JOIN movies m ON o.movies_id = m.id
		JOIN order_time ot ON o.order_time_id = ot.id
		JOIN order_location ol ON o.order_location_id = ol.id
		WHERE m.id = $1`
	rows, err := pkg.Database.Query(c, query, movie_id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var schedules []models.Schedule

	for rows.Next() {
		var schedule models.Schedule
		if err := rows.Scan(&schedule.Id, &schedule.Movie, &schedule.Date, &schedule.Time, &schedule.Location); err != nil {
			return nil, err
		}
		schedules = append(schedules, schedule)
	}

	return schedules, nil
}
