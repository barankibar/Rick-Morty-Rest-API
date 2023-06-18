package models

type User struct {
	UserName string `json:"username" validate:"required"`
	Password string `json:"password"`
}
