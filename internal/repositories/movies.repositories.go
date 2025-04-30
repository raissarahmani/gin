package repositories

import (
	"context"
	"main/internal/models"
	"main/pkg"
)

type MovieRepositories struct{}

var MovieRepo *MovieRepositories

func NewMovieRepository() {
	MovieRepo = &MovieRepositories{}
}

func (m *MovieRepositories) ShowAllMovies(c context.Context) ([]models.Movies, error) {
	query := `
		SELECT m.id, mi.poster, m.title, mg.genre_name
		FROM movies m
		JOIN movies_genre mg ON m.movies_genre_id = mg.id
		JOIN movies_image mi ON m.movies_image_id = mi.id`
	rows, err := pkg.Database.Query(c, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var movies []models.Movies
	for rows.Next() {
		var movie models.Movies
		if err := rows.Scan(&movie.Id, &movie.Image, &movie.Title, &movie.Genre); err != nil {
			return nil, err
		}
		movies = append(movies, movie)
	}

	return movies, nil
}

func (m *MovieRepositories) ShowMovieDetail(c context.Context, id int) (models.Movies, error) {
	query := `
		SELECT m.id, mi.poster, m.title, mg.genre_name, m.duration, m.release_date, m.director, m.casts, m.synopsis
		FROM movies m
		JOIN movies_genre mg ON m.movies_genre_id = mg.id
		JOIN movies_image mi ON m.movies_image_id = mi.id
		WHERE m.id = $1`

	var detail models.Movies
	err := pkg.Database.QueryRow(c, query, id).Scan(&detail.Id, &detail.Image, &detail.Title, &detail.Genre, &detail.Duration, &detail.Release_date, &detail.Director, &detail.Casts, &detail.Synopsis)

	if err != nil {
		return models.Movies{}, err
	}
	return detail, nil
}

func (m *MovieRepositories) FilterMoviesByTitle(c context.Context, title string) ([]models.Movies, error) {
	query := `
		SELECT m.id, mi.poster, m.title, mg.genre_name
		FROM movies m
		JOIN movies_genre mg ON m.movies_genre_id = mg.id
		JOIN movies_image mi ON m.movies_image_id = mi.id
		WHERE LOWER(m.title) = LOWER($1)`
	rows, err := pkg.Database.Query(c, query, title)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var movies []models.Movies
	for rows.Next() {
		var movie models.Movies
		if err := rows.Scan(&movie.Id, &movie.Image, &movie.Title, &movie.Genre); err != nil {
			return nil, err
		}
		movies = append(movies, movie)
	}

	return movies, nil
}

func (m *MovieRepositories) FilterMoviesByGenre(c context.Context, genreName string) ([]models.Movies, error) {
	query := `
		SELECT m.id, mi.poster, m.title, mg.genre_name
		FROM movies m
		JOIN movies_genre mg ON m.movies_genre_id = mg.id
		JOIN movies_image mi ON m.movies_image_id = mi.id
		WHERE LOWER(mg.genre_name) = LOWER($1)`
	rows, err := pkg.Database.Query(c, query, genreName)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var movies []models.Movies
	for rows.Next() {
		var movie models.Movies
		if err := rows.Scan(&movie.Id, &movie.Image, &movie.Title, &movie.Genre); err != nil {
			return nil, err
		}
		movies = append(movies, movie)
	}

	return movies, nil
}

func (m *MovieRepositories) FilterMoviesByTitleAndGenre(c context.Context, title, genre string) ([]models.Movies, error) {
	query := `
		SELECT m.id, mi.poster, m.title, mg.genre_name
		FROM movies m
		JOIN movies_genre mg ON m.movies_genre_id = mg.id
		JOIN movies_image mi ON m.movies_image_id = mi.id
		WHERE LOWER(m.title) = LOWER($1) AND LOWER(mg.genre_name) = LOWER($2)`
	var result []models.Movies
	rows, err := pkg.Database.Query(c, query, title, genre)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var movie models.Movies
		if err := rows.Scan(&movie.Id, &movie.Image, &movie.Title, &movie.Genre); err != nil {
			return nil, err
		}
		result = append(result, movie)
	}
	return result, nil
}

func (m *MovieRepositories) ShowUpcomingMovies(c context.Context) ([]models.Movies, error) {
	query := `
		SELECT m.id, mi.poster, m.title, mg.genre_name, m.release_date
		FROM movies m
		JOIN movies_genre mg ON m.movies_genre_id = mg.id
		JOIN movies_image mi ON m.movies_image_id = mi.id
		WHERE m.release_date > now()`
	rows, err := pkg.Database.Query(c, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var movies []models.Movies
	for rows.Next() {
		var movie models.Movies
		if err := rows.Scan(&movie.Id, &movie.Image, &movie.Title, &movie.Genre, &movie.Release_date); err != nil {
			return nil, err
		}

		movies = append(movies, movie)
	}

	return movies, nil
}

func (m *MovieRepositories) GetSchedulesByMovieID(c context.Context, movie_id int) ([]models.Schedule, error) {
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
