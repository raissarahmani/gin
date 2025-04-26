package models

type Seat struct {
	Id           int      `db:"id" json:"id"`
	Cinema       int      `db:"cinema_id" json:"cinema_id"`
	Seat_number  []string `db:"seat_number" json:"seat_number"`
	Available    int      `db:"seat_available" json:"available"`
	Jumlah_kursi int      `db:"qty" json:"qty"`
}
