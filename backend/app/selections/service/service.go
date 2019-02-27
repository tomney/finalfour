package service

import (
	"github.com/tomney/finalfour/backend/app/selections"
	"github.com/tomney/finalfour/backend/app/selections/repository"
)

// Interface implements the methods to interact with selections
type Interface interface {
	Create(selections.Selections) error
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
	//TODO  build this function
	return s.repo.Create(selections)
}
