package models

import (
	"mime/multipart"
	"time"
)

type Movies struct {
	Id           int       `db:"id" json:"id"`
	Image        string    `db:"movies_image_id" json:"poster"`
	Title        string    `db:"title" json:"title"`
	Genre        string    `db:"movies_genre_id" json:"genre"`
	Duration     int       `db:"duration" json:"duration,omitempty"`
	Release_date time.Time `db:"release_date" json:"release,omitempty"`
	Director     string    `db:"director" json:"director,omitempty"`
	Casts        []string  `db:"casts" json:"casts,omitempty"`
	Synopsis     string    `db:"synopsis" json:"synopsis,omitempty"`
}

type MoviesForm struct {
	Title        string                `form:"title"`
	Image        *multipart.FileHeader `form:"img"`
	Duration     int                   `form:"duration"`
	Release_date time.Time             `form:"release"`
	Director     string                `form:"director"`
	Synopsis     string                `form:"synopsis"`
}
