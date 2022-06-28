package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string `json:"email" gorm:"unique"`
	Nickname string `json:"nickname"`
	Provider string `json:"provider"`
	Minis    []Mini `json:"minis"`
}
