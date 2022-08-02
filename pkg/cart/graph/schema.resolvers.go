package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"log"

	"github.com/pawelkonikpl/ecommerce-api/pkg/cart/graph/generated"
	"github.com/pawelkonikpl/ecommerce-api/pkg/cart/graph/model"
)

// CreateCart is the resolver for the createCart field.
func (r *mutationResolver) CreateCart(ctx context.Context, input model.NewCart) (*model.Cart, error) {
	cart := new(model.Cart)
	user := new(model.User)
	user.ID = "1"
	cart.ID = "1"
	cart.CartUser = user
	log.Println("asdasdas")
	log.Println("CREATED cart")
	return cart, nil
}

// Cart is the resolver for the cart field.
func (r *queryResolver) Cart(ctx context.Context) ([]*model.Cart, error) {
	var cart []*model.Cart
	cart = append(cart, &model.Cart{
		ID: "1",
	})
	return cart, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
