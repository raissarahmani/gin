package models

import "time"

type Movies struct {
	Id           int       `db:"id" json:"id"`
	Image        string    `db:"movies_image_id" json:"poster"`
	Title        string    `db:"title" json:"movie_title"`
	Genre        string    `db:"movies_genre_id" json:"genre"`
	Duration     int       `db:"duration" json:"duration,omitempty"`
	Release_date time.Time `db:"release_date" json:"release,omitempty"`
	Director     string    `db:"director" json:"director,omitempty"`
	Casts        []string  `db:"casts" json:"casts,omitempty"`
	Synopsis     string    `db:"synopsis" json:"synopsis,omitempty"`
}

type Schedule struct {
	Id       int       `db:"id" json:"id"`
	Movie    string    `db:"movies_id" json:"title"`
	Date     time.Time `db:"book_date" json:"date"`
	Time     time.Time `db:"order_time_id" json:"time"`
	Location string    `db:"order_location_id" json:"location"`
}
