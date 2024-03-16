package models

type User struct {
	Email    string `gorm:"primary_key" json:"email"`
	Password string `json:"password"`
}
