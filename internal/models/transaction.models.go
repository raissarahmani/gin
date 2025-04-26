package models

type Transaction struct {
	Id             int    `db:"id" json:"id"`
	Seat_price     int    `db:"seat_price_id" json:"seat_price_id"`
	Fullname       string `db:"fullname" json:"fullname"`
	Email          string `db:"email" json:"email"`
	Phone          string `db:"phone" json:"phone"`
	Payment_method string `db:"payment_method" json:"payment_method"`
	Payment_done   bool   `db:"payment_done" json:"paid"`
}
