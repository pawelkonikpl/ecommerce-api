package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"log"

	"github.com/pawelkonikpl/ecommerce-api/pkg/users/graph/generated"
	"github.com/pawelkonikpl/ecommerce-api/pkg/users/graph/model"
)

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	user := new(model.User)
	user.Name = input.Name
	_, err := r.DB.Model(user).Insert()
	if err != nil {
		log.Printf("Cannot create user %s", err)
		return nil, err
	}
	return user, nil
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	var users []*model.User
	err := r.DB.Model(&users).Select()
	if err != nil {
		fmt.Printf("Cannot find users %s", err)
		return nil, err
	}
	return users, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
