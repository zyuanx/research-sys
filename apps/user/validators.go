package user

import (
	"github.com/go-playground/validator/v10"
)
type RegisterValidator struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
type LoginValidator struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (r *LoginValidator) GetError(err error) string {
	errors := err.(validator.ValidationErrors)
	for _, e := range errors {
		switch e.Field() {
		case "Username":
			switch e.ActualTag() {
			case "required":
				return "the username field is required"
			case "lt":
				return "the len of username must less than 5"
			}
		case "Password":
			switch e.ActualTag() {
			case "required":
				return "the password field is required"
			}
		}
	}

	return "param is error"
}
