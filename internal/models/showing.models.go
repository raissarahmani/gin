package models

import "time"

type Schedule struct {
	Id         int       `db:"id" json:"id,omitempty"`
	ScheduleID int       `db:"schedule_id" json:"scheduleID"`
	Date       time.Time `db:"book_date" json:"date,omitempty"`
	Time       time.Time `db:"book_time" json:"time,omitempty"`
	City       string    `db:"city" json:"city,omitempty"`
	Cinema     string    `db:"cinema_name" json:"cinema,omitempty"`
	Movie      string    `db:"title" json:"title,omitempty"`
	Seat       string    `db:"seat_number" json:"seat,omitempty"`
	Price      int       `db:"price" json:"price"`
}

type Seat struct {
	Id           int    `db:"id" json:"id,omitempty"`
	Seat_number  string `db:"seat" json:"seat,omitempty"`
	Is_available bool   `db:"is_available" json:"is_available"`
}
