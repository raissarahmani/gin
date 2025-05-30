package repositories

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"main/internal/models"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
)

type MovieRepositories struct {
	db  *pgxpool.Pool
	rdb *redis.Client
}

func NewMovieRepository(pg *pgxpool.Pool, rdc *redis.Client) *MovieRepositories {
	return &MovieRepositories{
		db:  pg,
		rdb: rdc,
	}
}

func (m *MovieRepositories) ShowAllMovies(c context.Context, limit, offset int) ([]models.Movies, error) {
	query := `
		SELECT m.id, mi.poster, m.title, string_agg(g.genre_name, ', ') as genre
		from movies m
		join movies_image mi on m.movies_image_id = mi.id
		JOIN LATERAL (
		  SELECT g.genre_name
		  FROM movies_genre mg 
		  JOIN genre g ON mg.genre_id = g.id
		  WHERE mg.movies_id = m.id
		  LIMIT 2) g ON true
		GROUP BY m.id, m.title, mi.poster
		ORDER BY m.id
		LIMIT $1 OFFSET $2`

	rows, err := m.db.Query(c, query, limit, offset)
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
		SELECT m.id, mi.poster, m.title, string_agg(g.genre_name, ', ') as genre, m.duration, m.release_date, m.director, m.casts, m.synopsis
		from movies m
		join movies_genre mg on mg.movies_id = m.id
		join genre g on mg.genre_id = g.id
		join movies_image mi on m.movies_image_id = mi.id
		WHERE m.id = $1
		group by m.id, m.title, mi.poster`

	var detail models.Movies
	err := m.db.QueryRow(c, query, id).Scan(&detail.Id, &detail.Image, &detail.Title, &detail.Genre, &detail.Duration, &detail.Release_date, &detail.Director, &detail.Casts, &detail.Synopsis)

	if err != nil {
		return models.Movies{}, err
	}
	return detail, nil
}

func (m *MovieRepositories) FilterMoviesByTitle(c context.Context, title string) ([]models.Movies, error) {
	query := `
		SELECT m.id, mi.poster, m.title, string_agg(g.genre_name, ', ') as genre
		from movies m
		join movies_genre mg on mg.movies_id = m.id
		join genre g on mg.genre_id = g.id
		join movies_image mi on m.movies_image_id = mi.id
		WHERE LOWER(m.title) = LOWER($1)
		group by m.id, m.title, mi.poster
		order by m.id`
	rows, err := m.db.Query(c, query, title)
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
		SELECT m.id, mi.poster, m.title, string_agg(g.genre_name, ', ') as genre
		from movies m
		join movies_genre mg on mg.movies_id = m.id
		join genre g on mg.genre_id = g.id
		join movies_image mi on m.movies_image_id = mi.id
		WHERE LOWER(g.genre_name) = LOWER($1)
		group by m.id, m.title, mi.poster
		order by m.id`
	rows, err := m.db.Query(c, query, genreName)
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
		SELECT m.id, mi.poster, m.title, string_agg(g.genre_name, ', ') as genre
		from movies m
		join movies_genre mg on mg.movies_id = m.id
		join genre g on mg.genre_id = g.id
		join movies_image mi on m.movies_image_id = mi.id
		WHERE LOWER(m.title) = LOWER($1) AND LOWER(g.genre_name) = LOWER($2)
		group by m.id, m.title, mi.poster
		order by m.id`
	var result []models.Movies
	rows, err := m.db.Query(c, query, title, genre)
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

func (m *MovieRepositories) ShowNowPlaying(c context.Context) ([]models.Movies, error) {
	redisKey := "movies:nowplaying"
	cache, err := m.rdb.Get(c, redisKey).Result()
	if err != nil {
		if err == redis.Nil {
			log.Printf("\nkey %s does not exist\n", redisKey)
		} else {
			log.Println("Redis is not working")
		}
	} else {
		var movies []models.Movies
		if err := json.Unmarshal([]byte(cache), &movies); err != nil {
			return nil, err
		}
		if len(movies) > 0 {
			return movies, nil
		}
	}

	query := `
		SELECT m.id, mi.poster, m.title, string_agg(g.genre_name, ', ') as genre
		from movies m
		join movies_image mi on m.movies_image_id = mi.id
		JOIN LATERAL (
		  SELECT g.genre_name
		  FROM movies_genre mg 
		  JOIN genre g ON mg.genre_id = g.id
		  WHERE mg.movies_id = m.id
		  LIMIT 2) g ON true
		WHERE m.id IN (6,7,8,9)
		group by m.id, m.title, mi.poster
		order by m.id`
	rows, err := m.db.Query(c, query)
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

	res, err := json.Marshal(movies)
	if err != nil {
		log.Println(err.Error())
	}
	if err := m.rdb.Set(c, redisKey, string(res), time.Minute*15).Err(); err != nil {
		log.Println(err.Error())
	}

	return movies, nil
}

func (m *MovieRepositories) ShowUpcomingMovies(c context.Context) ([]models.Movies, error) {
	// cek redis terlebih dahulu, jika ada nilainya maka gunakan nilai dari redis
	redisKey := "movies:upcoming"
	cache, err := m.rdb.Get(c, redisKey).Result()
	if err != nil {
		if err == redis.Nil {
			log.Printf("\nkey %s does not exist\n", redisKey)
		} else {
			log.Println("Redis is not working")
		}
	} else {
		var movies []models.Movies
		if err := json.Unmarshal([]byte(cache), &movies); err != nil {
			return nil, err
		}
		if len(movies) > 0 {
			return movies, nil
		}
	}

	query := `
		SELECT m.id, mi.poster, m.title, string_agg(g.genre_name, ', ') as genre, m.release_date
		from movies m
		join movies_image mi on m.movies_image_id = mi.id
		JOIN LATERAL (
		  SELECT g.genre_name
		  FROM movies_genre mg 
		  JOIN genre g ON mg.genre_id = g.id
		  WHERE mg.movies_id = m.id
		  LIMIT 2) g ON true
		WHERE m.release_date > now()
		group by m.id, m.title, mi.poster
		order by m.id`
	rows, err := m.db.Query(c, query)
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

	res, err := json.Marshal(movies)
	if err != nil {
		log.Println(err.Error())
	}
	if err := m.rdb.Set(c, redisKey, string(res), time.Minute*15).Err(); err != nil {
		log.Println(err.Error())
	}

	return movies, nil
}

func (m *MovieRepositories) AddNewMovie(c context.Context, movie models.Movies) error {
	query := `INSERT INTO movies (title, duration, release_date, director, synopsis) VALUES `
	values := []any{}
	query += `(`
	if movie.Title != "" {
		query += fmt.Sprintf("$%d,", len(values)+1)
		values = append(values, movie.Title)
	}
	if movie.Duration != 0 {
		query += fmt.Sprintf("$%d,", len(values)+1)
		values = append(values, movie.Duration)
	}
	if !movie.Release_date.IsZero() {
		query += fmt.Sprintf("$%d,", len(values)+1)
		values = append(values, movie.Release_date)
	}
	if movie.Director != "" {
		query += fmt.Sprintf("$%d,", len(values)+1)
		values = append(values, movie.Director)
	}
	if movie.Synopsis != "" {
		query += fmt.Sprintf("$%d,", len(values)+1)
		values = append(values, movie.Synopsis)
	}
	query += `)`

	log.Println("[DEBUG] qry: ", query)

	_, err := m.db.Exec(c, query, values...)
	return err
}

func (m *MovieRepositories) EditMovie(c context.Context, id int, movie models.Movies) error {
	if movie.Image != "" {
		query := `
			UPDATE movies
			SET movies_image_id=$1, title=$2, duration=$3, release_date=$4, director=$5, synopsis=$6
			WHERE id=$7`
		_, err := m.db.Exec(c, query, movie.Image, movie.Title, movie.Duration, movie.Release_date, movie.Director, movie.Synopsis, id)
		return err
	} else {
		query := `
			UPDATE movies
			SET title=$1, duration=$2, release_date=$3, director=$4, synopsis=$5
			WHERE id=$6`
		_, err := m.db.Exec(c, query, movie.Title, movie.Duration, movie.Release_date, movie.Director, movie.Synopsis, id)
		return err
	}
}

func (m *MovieRepositories) DeleteMovie(c context.Context, id int) error {
	query := `DELETE FROM movies WHERE id=$1`
	_, err := m.db.Exec(c, query, id)
	return err
}
