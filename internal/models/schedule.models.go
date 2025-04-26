package models

type Schedule struct {
	Id       int    `db:"id" json:"id"`
	User     int    `db:"users_id" json:"user_id"`
	Movie    int    `db:"movies_id" json:"movie_id"`
	Date     string `db:"book_date" json:"date"`
	Time     int    `db:"order_time_id" json:"time"`
	Location int    `db:"order_location_id" json:"location"`
}
