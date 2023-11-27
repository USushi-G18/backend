package models

import "gopkg.in/guregu/null.v4"

type Menu string

const (
	MenuCarte  Menu = "Carte"
	MenuLunch  Menu = "Lunch"
	MenuDinner Menu = "Dinner"
)

type Category struct {
	ID   int
	Name string
}

type Image struct {
	ID    int
	Image string
}

type Product struct {
	ID          int
	Name        string
	Price       string
	Category    int
	Menu        Menu
	Description null.String
	ImageID     null.Int `db:"image_id"`
	OrderLimit  null.Int `db:"order_limit"`
	Pieces      int
}

type Allergen struct {
	ID   int
	Name string
}

type Ingredient struct {
	ID       int
	Name     string
	Allergen null.Int
}
