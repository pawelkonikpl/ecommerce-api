package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/pawelkonikpl/ecommerce-api/cart/graph/generated"
	"github.com/pawelkonikpl/ecommerce-api/cart/graph/model"
)

// Cart is the resolver for the cart field.
func (r *queryResolver) Cart(ctx context.Context) ([]*model.Cart, error) {
	panic(fmt.Errorf("not implemented"))
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
