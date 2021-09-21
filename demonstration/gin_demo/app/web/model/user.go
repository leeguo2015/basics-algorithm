package model

import (
	"gorm.io/gorm"
)

var db *gorm.DB

type (
	todoModel struct {
		Model
		Title     string `json:"title"`
		Completed int    `json:"completed"`
	}

	fmtTodo struct {
		ID        uint   `json:"id"`
		Title     string `json:"title"`
		Completed bool   `json:"completed"`
	}
)

