package models

import "time"

const books = "books"

type Book struct {
	ID              int64
	Name            string
	Author          string
	PublicationDate time.Time
}

func (Book) TableName() string {
	return books
}
