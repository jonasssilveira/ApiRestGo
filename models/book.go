package models

import (
	uuid "github.com/satori/go.uuid"
)

type Book struct {
	Base
	Name        string  `json:"name"`
	Author      string  `json:"author"`
	Description string  `json:"description"`
	MediumPrice float64 `json:"medium_price"`
	ImageUrl    string  `json:"img_url"`
}

func NewBook() *Book {
	book := Book{Base: Base{Id: uuid.NewV4()}}
	return &book
}
