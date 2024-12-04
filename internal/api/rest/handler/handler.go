package handler

import (
	"context"

	"ingrytech/internal/models"
)

type App interface {
	CreateBook(ctx context.Context, book models.Book) (models.Book, error)
	Books(ctx context.Context) ([]models.Book, error)
	Book(ctx context.Context, id int64) (models.Book, error)
	UpdateBook(ctx context.Context, book models.Book) (models.Book, error)
	DeleteBook(ctx context.Context, id int64) error
}

type Handler struct {
	app App
}

func New(app App) *Handler {
	return &Handler{app: app}
}
