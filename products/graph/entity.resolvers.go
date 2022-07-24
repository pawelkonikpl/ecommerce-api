package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"log"

	"github.com/pawelkonikpl/ecommmerce-api/products/graph/generated"
	"github.com/pawelkonikpl/ecommmerce-api/products/graph/model"
)

// FindProductByID is the resolver for the findProductByID field.
func (r *entityResolver) FindProductByID(ctx context.Context, id string) (*model.Product, error) {
	product := new(model.Product)
	err := r.DB.Model(product).Where("id = ?", id).Select()
	if err != nil {
		log.Printf("Cannot find product %s", err)
		return nil, err
	}
	return product, nil
}

// FindUserByID is the resolver for the findUserByID field.
func (r *entityResolver) FindUserByID(ctx context.Context, id string) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

// Entity returns generated.EntityResolver implementation.
func (r *Resolver) Entity() generated.EntityResolver { return &entityResolver{r} }

type entityResolver struct{ *Resolver }
