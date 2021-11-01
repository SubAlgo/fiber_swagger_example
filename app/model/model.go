package model

import (
	"time"
)

// Error example
type APIError struct {
	ErrorCode    int
	ErrorMessage string
	CreatedAt    time.Time
}

type Users struct {
	Users []User `json:"users"`
}

type User struct {
	UserForCreate
	ID   int    `json:"id" example:"1"` // User ID
	Hash string `swaggerignore:"true"`
}

type UserForCreate struct {
	Username string `json:"username" binding:"required" example:"username" maxLength:"255"` // user Username
	UserForUpdate
}

type UserForUpdate struct {
	Firstname string `json:"firstname" binding:"required" example:"Firstname" maxLength:"255"` // User Firstname
	Lastname  string `json:"lastname" binding:"required" example:"Lastname" maxLength:"255"`   // User Lastname
	Email     string `json:"email" binding:"required" example:"abc@gmail.com" maxLength:"255"` // User E-mail
	Gender    string `json:"gender" example:"male" enums:"male,female"`                        // User Gender
}

type Response struct {
	Code   int    `json:code`   // Response Code
	Status string `json:status` // Response Status
}
