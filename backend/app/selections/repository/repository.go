package repository

import (
	"log"

	"github.com/tomney/finalfour/backend/app/selections"
)

// Interface implements the methods to store data/ mutate stored data for selections
type Interface interface {
	Create(selections selections.Selections) error
	delete(email string) error
}

// Repository handles storing data/ stored data for selections
type Repository struct {
	sqlClient string
}

// NewRepository returns a new repository instance
func NewRepository(sqlClient string) *Repository {
	return &Repository{sqlClient: sqlClient}
}

// Create creates an entry for selections
func (r *Repository) Create(selections selections.Selections) error {
	//TODO  build this function
	log.Printf("Creating the selections entry.")
	return nil
}

func (r *Repository) delete(email string) error {
	//TODO build this function, should do a soft delete
	return nil
}
