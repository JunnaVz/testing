package models

import "github.com/google/uuid"

type User struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Surname     string    `json:"surname"`
	Address     string    `json:"address"`
	PhoneNumber string    `json:"phoneNumber"`
	Email       string    `json:"email"`
	Password    string    `json:"password"`
}
