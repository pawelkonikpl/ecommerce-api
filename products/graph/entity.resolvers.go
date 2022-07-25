package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"log"

	"github.com/pawelkonikpl/ecommerce-api/products/graph/generated"
	"github.com/pawelkonikpl/ecommerce-api/products/graph/model"
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

// Entity returns generated.EntityResolver implementation.
func (r *Resolver) Entity() generated.EntityResolver { return &entityResolver{r} }

type entityResolver struct{ *Resolver }
