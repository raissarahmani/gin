package models

import "mime/multipart"

type Profile struct {
	Id         int    `db:"id" json:"id,omitempty"`
	User       int    `db:"users_id" json:"user_id"`
	Image      string `db:"profile_image" json:"profile_image"`
	First_name string `db:"first_name" json:"first_name"`
	Last_name  string `db:"last_name" json:"last_name"`
	Email      string `db:"email" json:"email"`
	Phone      string `db:"phone" json:"phone"`
	Password   string `db:"password" json:"-"`
}

type ProfileForm struct {
	Image *multipart.FileHeader `form:"img"`
}
