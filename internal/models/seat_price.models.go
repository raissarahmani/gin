package models

type SeatPrice struct {
	Id    int `db:"id" json:"id"`
	Seat  int `db:"seat_id" json:"seat_id"`
	Total int `db:"total_price" json:"total"`
}
