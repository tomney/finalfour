package repository

import (
	"database/sql"
	"log"

	"github.com/tomney/finalfour/backend/app/selections"
	"github.com/tomney/finalfour/backend/app/team"
)

// Interface implements the methods to store data/ mutate stored data for selections
type Interface interface {
	Create(selections selections.Selections) error
	delete(email string) error
	Get(email string) (*selections.Selections, error)
	List() error
}

// Repository handles storing data/ stored data for selections
type Repository struct {
	db sql.DB
}

// NewRepository returns a new repository instance
func NewRepository(sqlDB sql.DB) *Repository {
	return &Repository{db: sqlDB}
}

// Create creates an entry for selections
func (r *Repository) Create(selections selections.Selections) error {
	existingSelections, err := r.Get(selections.Email)
	if err != nil {
		log.Printf("An error occurred trying to get the existing selections")
		return err
	}
	if existingSelections != nil {
		err := r.delete(selections.Email)
		if err != nil {
			log.Printf("Unable to create new selections as an error occurred deleting old selections")
			return err
		}
	}

	insertStatement := `
	INSERT INTO selections (email, first_pick, second_pick, third_pick, fourth_pick, created)
		VALUES(?, ?, ?, ?, ?, DATE(NOW())
	`
	stmt, err := r.db.Prepare(insertStatement)
	if err != nil {
		log.Printf("Unable to prepare the statement to insert a selection")
		return err
	}
	_, err = stmt.Exec(selections.Email, selections.Teams[0], selections.Teams[1], selections.Teams[2], selections.Teams[3])
	if err != nil {
		log.Printf("Error occured while trying to insert a selections entry")
		return err
	}

	return nil
}

func (r *Repository) delete(email string) error {
	deleteStatement := `
		DELETE FROM selections
		WHERE email = ?
	`
	stmt, err := r.db.Prepare(deleteStatement)
	if err != nil {
		log.Printf("Unable to prepare the statement to delete a selections entry")
		return err
	}
	_, err = stmt.Exec(email)
	if err != nil {
		log.Printf("Error occured while trying to delete a selections entry")
		return err
	}

	return nil
}

// Get gets the existing selections entry for a given email if it exists
func (r *Repository) Get(email string) (*selections.Selections, error) {
	queryStatement := `
		SELECT email, first_pick, second_pick, third_pick, fourth_pick, created
		FROM selections
		WHERE email = ?
	`
	var first, second, third, fourth, created string
	err := r.db.QueryRow(queryStatement, email).Scan(&email, &first, &second, &third, &fourth, &created)
	if err != nil {
		log.Printf("Unable to get selections for email: %s", email)
		return nil, err
	}
	return &selections.Selections{
		Email: email,
		Teams: []team.Team{{ID: first}, {ID: second}, {ID: third}, {ID: fourth}},
	}, nil
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
