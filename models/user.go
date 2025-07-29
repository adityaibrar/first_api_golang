package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `json : "username" gorm : "unique"`
	Password string `json: "password"`
	Email    string `json: "email" gorm: "unique"`
}

type LoginRequest struct {
	Username string `json: "username"`
	Password string `json: "password"`
}

type RegisterRequest struct {
	Username string `json: "username"`
	Password string `json: "password"`
	Email    string `json: "email"`
}
