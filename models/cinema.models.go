package models

type Cinema struct {
	Id       int    `db:"id" json:"id"`
	Schedule int    `db:"order_id" json:"schedule_id"`
	Cinema   int    `db:"cinema_name_id" json:"cinema_id"`
	Studio   string `db:"studio_type" json:"studio"`
	Price    int    `db:"price" json:"price"`
}
