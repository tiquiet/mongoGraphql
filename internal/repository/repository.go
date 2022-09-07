package repository

import (
	"context"
	"github.com/tiquiet/mongoGraphql/internal/custom_models"
)

type Repository interface {
	CreateBook(ctx context.Context, input custom_models.NewBook) (string, error)
	GetBook(ctx context.Context, id string) (*custom_models.Book, error)
	GetAllBooks(ctx context.Context) ([]*custom_models.Book, error)
	UpdateBook(ctx context.Context, id string, input custom_models.UpdateBook) (bool, error)
	DeleteBook(ctx context.Context, id string) (bool, error)
}
