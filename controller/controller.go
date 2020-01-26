package controller

import (
	"go-restful-product-search-webapp/database"
)

// Controller is the main controller struct upon which other controller functions build.
type Controller struct {
	database.DB
}

// New creates a new controller struct and sets the appropriate data access objects.
func New(db *database.DB) *Controller {
	return &Controller{
		DB: *db,
	}
}
