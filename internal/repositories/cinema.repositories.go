package repositories

import (
	"context"
	"main/internal/models"
	"main/pkg"
)

type CinemaRepositories struct{}

var CinemaRepo *CinemaRepositories

func NewCinemaRepository() {
	CinemaRepo = &CinemaRepositories{}
}

func (c *CinemaRepositories) GetCinemaBySchedule(ctx context.Context, schedule_id int) ([]models.Cinema, error) {
	query := `
		select c.id, ol.location, cn.cinema_name, c.studio_type, c.price
		from cinema c
		join "order" o on c.order_id = o.id
		join order_location ol on o.order_location_id = ol.id
		join cinema_name cn on c.cinema_name_id = cn.id
		WHERE o.id = $1`
	rows, err := pkg.Database.Query(ctx, query, schedule_id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cinemas []models.Cinema

	for rows.Next() {
		var cinema models.Cinema
		if err := rows.Scan(&cinema.Id, &cinema.Location, &cinema.Cinema, &cinema.Studio, &cinema.Price); err != nil {
			return nil, err
		}
		cinemas = append(cinemas, cinema)
	}

	return cinemas, nil
}
