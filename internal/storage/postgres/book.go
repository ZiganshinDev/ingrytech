package postgres

import (
	"context"

	"ingrytech/internal/models"
)

func (db *DB) CreateBook(ctx context.Context, book models.Book) (models.Book, error) {
	err := db.WithContext(ctx).Model(&models.Book{}).Create(&book).Error
	if err != nil {
		return models.Book{}, err
	}

	return book, nil
}

func (db *DB) Books(ctx context.Context) ([]models.Book, error) {
	var books []models.Book

	err := db.WithContext(ctx).Model(&models.Book{}).Find(&books).Error
	if err != nil {
		return nil, err
	}

	return books, nil
}

func (db *DB) Book(ctx context.Context, id int64) (models.Book, error) {
	var book models.Book

	err := db.WithContext(ctx).Model(&models.Book{}).Where("id = ?", id).Find(&book).Error
	if err != nil {
		return models.Book{}, err
	}

	return book, nil
}

func (db *DB) UpdateBook(ctx context.Context, book models.Book) (models.Book, error) {
	err := db.WithContext(ctx).Updates(&book).Error
	if err != nil {
		return models.Book{}, err
	}

	return book, nil
}

func (db *DB) DeleteBook(ctx context.Context, id int64) error {
	return db.WithContext(ctx).Delete(&models.Book{}, id).Error
}
