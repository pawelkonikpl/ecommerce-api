package main

import (
	"github.com/pawelkonikpl/ecommerce-api/pkg/cart"
	"github.com/pawelkonikpl/ecommerce-api/pkg/products"
	"github.com/pawelkonikpl/ecommerce-api/pkg/users"
)

func main() {
	go products.Run()
	users.Run()
	cart.Run()
}
