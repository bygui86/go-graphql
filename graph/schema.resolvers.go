package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/bygui86/go-graphql/graph/generated"
	"github.com/bygui86/go-graphql/graph/model"
)

func (r *mutationResolver) CreateProduct(ctx context.Context, input model.NewProduct) (*model.Product, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateProduct(ctx context.Context, productID int, input model.NewProduct) (*model.Product, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteProduct(ctx context.Context, productID int) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Products(ctx context.Context) ([]*model.Product, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Product(ctx context.Context) (*model.Product, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
