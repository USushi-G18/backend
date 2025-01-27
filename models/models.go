package models

import (
	"time"

	"gopkg.in/guregu/null.v4"
)

type Menu string
type UserType string
type CommandStatus string

const (
	MenuCarte  Menu = "Carte"
	MenuLunch  Menu = "Lunch"
	MenuDinner Menu = "Dinner"
)

const (
	UserAdmin    UserType = "Admin"
	UserClient   UserType = "Client"
	UserEmployee UserType = "Employee"
)

const (
	CommandOrdered   CommandStatus = "Ordered"
	CommandPreparing CommandStatus = "Preparing"
	CommandPrepared  CommandStatus = "Prepared"
	CommandDelivered CommandStatus = "Delivered"
)

type Category struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Image struct {
	ID    int    `json:"id"`
	Image string `json:"image"`
}

type Plate struct {
	ID          int         `json:"id"`
	Name        string      `json:"name"`
	Price       string      `json:"price"`
	CategoryID  int         `json:"categoryID" db:"category_id"`
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
	ID         int      `json:"id"`
	Name       string   `json:"name"`
	AllergenID null.Int `json:"allergenID" db:"allergen_id"`
}

type Command struct {
	SessionID int           `json:"sessionID" db:"session_id"`
	PlateID   int           `json:"plateID" db:"plate_id"`
	At        time.Time     `json:"at"`
	Quantity  int           `json:"quantity"`
	Status    CommandStatus `json:"status"`
}

type Session struct {
	ID          int       `json:"id"`
	StartAt     time.Time `json:"startAt" db:"start_at"`
	EndAt       null.Time `json:"endAt" db:"end_at"`
	TableNumber int       `json:"tableNumber" db:"table_number"`
	Menu        Menu      `json:"menu"`
	Seating     int       `json:"seating"`
}

type ReturningID struct {
	ID int `json:"id"`
}
