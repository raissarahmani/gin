package models

type Response struct {
	Msg  string `json:"message"`
	Data any    `json:"data"`
}
