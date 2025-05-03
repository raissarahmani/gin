package models

import "time"

type Schedule struct {
	Id     int       `db:"id" json:"id"`
	Date   time.Time `db:"schedule_id" json:"date,omitempty"`
	Time   time.Time `db:"schedule_id" json:"time,omitempty"`
	City   string    `db:"city_id" json:"city,omitempty"`
	Cinema string    `db:"cinema_id" json:"cinema,omitempty"`
	Movie  string    `db:"movies_id" json:"title,omitempty"`
	Seat   string    `db:"seat_id" json:"seat,omitempty"`
}
