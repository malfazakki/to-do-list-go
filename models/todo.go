package models

import "gorm.io/gorm"

type Todo struct {
	gorm.Model
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}
