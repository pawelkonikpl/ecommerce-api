package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/pawelkonikpl/ecommerce-api/pkg/users/graph/generated"
	"github.com/pawelkonikpl/ecommerce-api/pkg/users/graph/model"
)

// FindUserByID is the resolver for the findUserByID field.
func (r *entityResolver) FindUserByID(ctx context.Context, id string) (*model.User, error) {
	user := new(model.User)
	err := r.DB.Model(user).Where("id = ?", id).Select()
	if err != nil {
		fmt.Printf("Cannot find user %s", err)
		return nil, err
	}
	return user, nil
}

// Entity returns generated.EntityResolver implementation.
func (r *Resolver) Entity() generated.EntityResolver { return &entityResolver{r} }

type entityResolver struct{ *Resolver }
