package service

import (
	"log"

	"github.com/tomney/finalfour/backend/app/selections"
	"github.com/tomney/finalfour/backend/app/selections/repository"
)

// Interface implements the methods to interact with selections
type Interface interface {
	Create(selections.Selections) error
	List() ([]selections.Selections, error)
}

// Service handles the collection and alteration of selections
type Service struct {
	repo repository.Interface
}

// NewService returns a new service instance
func NewService(repo repository.Interface) *Service {
	return &Service{repo: repo}
}

// Create creates a selections entry
func (s *Service) Create(selections selections.Selections) error {
	teamIDs, err := s.repo.Get(selections.Email)
	if err != nil {
		log.Printf("An error occurred trying to get the existing selections")
		return err
	}
	if teamIDs != nil {
		err := s.repo.Delete(selections.Email)
		if err != nil {
			log.Printf("Unable to create new selections as an error occurred deleting old selections")
			return err
		}
	}
	return s.repo.Create(selections)
}

// List lists the selections
func (s *Service) List() ([]selections.Selections, error) {
	//Hardcode selection responses for now
	selection1 := selections.Stub
	selection2 := selections.Stub
	selections := []selections.Selections{selection1, selection2}

	//TODO implement the service layer (will probably combine teams and selection data)
	_, _ = s.repo.List()
	return selections, nil
}
