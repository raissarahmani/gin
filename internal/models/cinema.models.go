package models

type Cinema struct {
	Id       int    `db:"id" json:"id"`
	Location string `db:"order_id" json:"city"`
	Cinema   string `db:"cinema_name_id" json:"cinema_name"`
	Studio   string `db:"studio_type" json:"studio"`
	Price    int    `db:"price" json:"price"`
}
