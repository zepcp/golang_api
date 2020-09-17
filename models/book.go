package models

import (
	"github.com/jinzhu/gorm"
)

// Book - Book Info
type Book struct {
	gorm.Model

	Isbn   string `json:"isbn"`
	Title  string `json:"title"`
	Author string `json:"author"`
}
