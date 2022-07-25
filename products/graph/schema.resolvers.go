package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"log"

	"github.com/pawelkonikpl/ecommerce-api/products/graph/generated"
	"github.com/pawelkonikpl/ecommerce-api/products/graph/model"
)

// CreateProduct is the resolver for the createProduct field.
func (r *mutationResolver) CreateProduct(ctx context.Context, input model.NewProduct) (*model.Product, error) {
	product := new(model.Product)
	product.Name = input.Name
	product.Sku = input.Sku
	product.CreatedAt = "1"
	_, err := r.DB.Model(product).Insert()
	if err != nil {
		log.Printf("Cannot create user %s", err)
		return nil, err
	}
	return product, nil
}

// Products is the resolver for the products field.
func (r *queryResolver) Products(ctx context.Context) ([]*model.Product, error) {
	var products []*model.Product
	err := r.DB.Model(&products).Select()
	if err != nil {
		log.Printf("Cannot find products %s", err)
		return nil, err
	}
	return products, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
