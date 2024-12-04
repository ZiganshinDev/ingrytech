package postgres

import (
	"go.uber.org/zap"
	"golang.org/x/net/context"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"moul.io/zapgorm2"
)

type DB struct {
	*gorm.DB
}

func New(ctx context.Context, uri string, logger *zap.SugaredLogger) (*DB, error) {
	l := zapgorm2.New(logger.Desugar())
	l.SetAsDefault()

	db, err := gorm.Open(postgres.Open(uri), &gorm.Config{
		Logger: l,
	})
	if err != nil {
		return nil, err
	}

	return &DB{db}, nil
}
