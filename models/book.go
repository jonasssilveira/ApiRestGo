package models

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
	"time"
)

type Book struct {
	Id          uuid.UUID      `json:"id" gorm:"type:uuid;primaryKey;id"`
	CreatedAt   time.Time      `json:"created"`
	UpdatedAt   time.Time      `json:"updated"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted"`
	Name        string         `json:"name"`
	Author      string         `json:"author"`
	Description string         `json:"description"`
	MediumPrice float64        `json:"medium_price"`
	ImageUrl    string         `json:"img_url"`
}

func NewBook() *Book {
	book := Book{Id: uuid.NewV4()}
	return &book
}
