package models

import "time"

type Schedule struct {
	Id         int       `db:"id" json:"id"`
	Date       time.Time `db:"book_date" json:"date,omitempty"`
	Time       time.Time `db:"book_time" json:"time,omitempty"`
	City       string    `db:"city" json:"city,omitempty"`
	Cinema     string    `db:"cinema_name" json:"cinema,omitempty"`
	Movie      string    `db:"title" json:"title,omitempty"`
	Seat       string    `db:"seat_number" json:"seat,omitempty"`
	MovieID    int       `json:"movie_id,omitempty"`
	CityID     int       `json:"city_id,omitempty"`
	CinemaID   int       `json:"cinema_id,omitempty"`
	ScheduleID int       `json:"schedule_id,omitempty"`
	SeatID     int       `json:"seat_id,omitempty"`
}
