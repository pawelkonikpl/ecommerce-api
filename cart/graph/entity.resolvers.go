package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/pawelkonikpl/ecommerce-api/cart/graph/generated"
	"github.com/pawelkonikpl/ecommerce-api/cart/graph/model"
)

// FindCartByID is the resolver for the findCartByID field.
func (r *entityResolver) FindCartByID(ctx context.Context, id string) (*model.Cart, error) {
	cart := new(model.Cart)
	cart.ID = "1"
	cart.CartUser.ID = "1"
	return cart, nil
}

// FindUserByID is the resolver for the findUserByID field.
func (r *entityResolver) FindUserByID(ctx context.Context, id string) (*model.User, error) {
	user := new(model.User)
	user.ID = "1"
	return user, nil
}

// Entity returns generated.EntityResolver implementation.
func (r *Resolver) Entity() generated.EntityResolver { return &entityResolver{r} }

type entityResolver struct{ *Resolver }
