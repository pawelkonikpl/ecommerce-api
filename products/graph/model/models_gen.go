// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type NewProduct struct {
	UserID string `json:"userId"`
	Name   string `json:"name"`
	Sku    string `json:"sku"`
}

type Product struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Sku       string `json:"sku"`
	CreatedAt string `json:"createdAt"`
}

func (Product) IsEntity() {}

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (User) IsEntity() {}
