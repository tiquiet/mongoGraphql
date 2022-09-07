package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/tiquiet/mongoGraphql/graph/generated"
	"github.com/tiquiet/mongoGraphql/internal/custom_models"
)

// CreateBook is the resolver for the createBook field.
func (r *mutationResolver) CreateBook(ctx context.Context, input custom_models.NewBook) (string, error) {
	return r.Repository.CreateBook(ctx, input)
}

// UpdateBook is the resolver for the updateBook field.
func (r *mutationResolver) UpdateBook(ctx context.Context, id string, input custom_models.UpdateBook) (bool, error) {
	return r.Repository.UpdateBook(ctx, id, input)
}

// DeleteBook is the resolver for the deleteBook field.
func (r *mutationResolver) DeleteBook(ctx context.Context, id string) (bool, error) {
	return r.Repository.DeleteBook(ctx, id)
}

// GetAllBooks is the resolver for the getAllBooks field.
func (r *queryResolver) GetAllBooks(ctx context.Context) ([]*custom_models.Book, error) {
	return r.Repository.GetAllBooks(ctx)
}

// GetBook is the resolver for the getBook field.
func (r *queryResolver) GetBook(ctx context.Context, id string) (*custom_models.Book, error) {
	return r.Repository.GetBook(ctx, id)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
