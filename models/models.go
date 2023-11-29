package models

import "gopkg.in/guregu/null.v4"

type Menu string
type UserType string

const (
	MenuCarte  Menu = "Carte"
	MenuLunch  Menu = "Lunch"
	MenuDinner Menu = "Dinner"
)

const (
	UserAdmin  UserType = "Admin"
	UserClient UserType = "Client"
)

type Category struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Image struct {
	ID    int    `json:"id"`
	Image string `json:"image"`
}

type Product struct {
	ID          int         `json:"id"`
	Name        string      `json:"name"`
	Price       string      `json:"price"`
	Category    int         `json:"category"`
	Menu        Menu        `json:"menu"`
	Description null.String `json:"description"`
	ImageID     null.Int    `json:"imageID" db:"image_id"`
	OrderLimit  null.Int    `json:"orderLimit" db:"order_limit"`
	Pieces      int         `json:"pieces"`
}

type Allergen struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Ingredient struct {
	ID       int      `json:"id"`
	Name     string   `json:"name"`
	Allergen null.Int `json:"allergen"`
}
