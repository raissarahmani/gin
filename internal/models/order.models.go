package models

import "time"

type OrderRequest struct {
	MovieID         int    `json:"movie_id"`
	CityID          int    `json:"city_id"`
	CinemaID        int    `json:"cinema_id"`
	ScheduleID      int    `json:"schedule_id"`
	SeatIDs         []int  `json:"seat_ids"`
	Fullname        string `json:"fullname"`
	Email           string `json:"email"`
	Phone           string `json:"phone"`
	PaymentMethodID int    `json:"payment_method_id"`
	Total           int    `json:"total"`
}

type OrderHistory struct {
	TransactionID int       `json:"transaction_id"`
	BookDate      time.Time `json:"book_date"`
	BookTime      time.Time `json:"book_time"`
	City          string    `json:"city"`
	CinemaName    string    `json:"cinema_name"`
	MovieTitle    string    `json:"movie_title"`
	SeatNumber    string    `json:"seat_number"`
	Total         int       `json:"total"`
	PaymentDone   bool      `json:"payment_done"`
}
