package app

import (
	"context"

	"ingrytech/internal/models"
)

type DB interface {
	CreateBook(ctx context.Context, book models.Book) (models.Book, error)
	Books(ctx context.Context) ([]models.Book, error)
	Book(ctx context.Context, id int64) (models.Book, error)
	UpdateBook(ctx context.Context, book models.Book) (models.Book, error)
	DeleteBook(ctx context.Context, id int64) error
}

type App struct {
	db DB
}

func New(db DB) *App {
	return &App{db: db}
}
