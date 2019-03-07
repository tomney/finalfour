package repository

import (
	"database/sql"
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
	// sqlClient cloudsql.Client
	db sql.DB
}

// NewRepository returns a new repository instance
func NewRepository(sqlDB sql.DB) *Repository {
	return &Repository{db: sqlDB}
}

// Create creates an entry for selections
func (r *Repository) Create(selections selections.Selections) error {
	//TODO  build this function
	log.Printf("Creating the selections entry.")
	return nil
}

// List gets the selection entries
func (r *Repository) List() error {
	// test that the cloud sql instance works
	rows, err := r.db.Query("SELECT * FROM selections;")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var (
			email   string
			first   string
			second  string
			third   string
			fourth  string
			created string
		)
		if err := rows.Scan(&email, &first, &second, &third, &fourth, &created); err != nil {
			log.Fatal(err)
		}
		log.Printf("email: %s", email)
		log.Printf("first: %s", first)
		log.Printf("second: %s", second)
		log.Printf("third: %s", third)
		log.Printf("fourth: %s", fourth)
		log.Printf("created: %s", created)
	}
	return nil
}

func (r *Repository) delete(email string) error {
	//TODO build this function, should do a soft delete
	return nil
}
