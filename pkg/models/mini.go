package models

import "gorm.io/gorm"

type Mini struct {
	gorm.Model
	Url    string `json:"url"`
	UserID uint   `json:"userId"`
	Mini   string `json:"mini" gorm:"-:all"`
}
