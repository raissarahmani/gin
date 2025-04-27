package models

import "time"

type Schedule struct {
	Id int `db:"id" json:"id"`
	// User     int       `db:"users_id" json:"user_id"`
	Movie    string    `db:"movies_id" json:"title"`
	Date     time.Time `db:"book_date" json:"date"`
	Time     time.Time `db:"order_time_id" json:"time"`
	Location string    `db:"order_location_id" json:"location"`
}
