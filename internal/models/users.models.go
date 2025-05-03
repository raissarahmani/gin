package models

type Users struct {
	Id       int    `db:"id" json:"-"`
	Email    string `db:"email" json:"email" binding:"required,email"`
	Password string `db:"password" json:"password,omitempty" binding:"required,min=8"`
	Role     string `db:"role" json:"role,omitempty"`
}
