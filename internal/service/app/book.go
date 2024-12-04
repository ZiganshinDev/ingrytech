package app

import (
	"context"

	"ingrytech/internal/models"
	"ingrytech/internal/svcerr"
)

func (a *App) CreateBook(ctx context.Context, req models.Book) (models.Book, error) {
	book, err := a.db.CreateBook(ctx, req)
	if err != nil {
		return models.Book{}, svcerr.NewErr(svcerr.ErrInternalError, err.Error())
	}

	return book, nil
}

func (a *App) Books(ctx context.Context) ([]models.Book, error) {
	books, err := a.db.Books(ctx)
	if err != nil {
		return nil, svcerr.NewErr(svcerr.ErrInternalError, err.Error())
	}

	return books, nil
}

func (a *App) Book(ctx context.Context, id int64) (models.Book, error) {
	book, err := a.db.Book(ctx, id)
	if err != nil {
		return models.Book{}, svcerr.NewErr(svcerr.ErrInternalError, err.Error())
	}

	if book.ID == 0 {
		return models.Book{}, svcerr.ErrNotFound
	}

	return book, nil
}

func (a *App) UpdateBook(ctx context.Context, req models.Book) (models.Book, error) {
	book, err := a.db.UpdateBook(ctx, req)
	if err != nil {
		return models.Book{}, svcerr.NewErr(svcerr.ErrInternalError, err.Error())
	}

	return book, nil
}

func (a *App) DeleteBook(ctx context.Context, id int64) error {
	if err := a.db.DeleteBook(ctx, id); err != nil {
		return svcerr.NewErr(svcerr.ErrInternalError, err.Error())
	}

	return nil
}
