package models

import "mime/multipart"

type Profile struct {
	User       int    `db:"users_id" json:"-"`
	Email      string `db:"email" json:"email"`
	Image      string `db:"profile_image" json:"profile_image"`
	First_name string `db:"first_name" json:"first_name"`
	Last_name  string `db:"last_name" json:"last_name"`
	Phone      string `db:"phone" json:"phone"`
	Password   string `db:"password" json:"-"`
}

type ProfileForm struct {
	First_name string                `form:"first_name"`
	Last_name  string                `form:"last_name"`
	Phone      string                `form:"phone"`
	Email      string                `form:"email" binding:"email"`
	Image      *multipart.FileHeader `form:"img"`
}
