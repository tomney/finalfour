package service

import (
	"log"

	model "github.com/tomney/finalfour/backend/app/selections"
	"github.com/tomney/finalfour/backend/app/selections/repository"
	teamservice "github.com/tomney/finalfour/backend/app/team/service"
)

// Interface implements the methods to interact with selections
type Interface interface {
	Create(model.Selections) error
	List() ([]model.Selections, error)
}

// Service handles the collection and alteration of selections
type Service struct {
	repo repository.Interface
	team teamservice.Interface
}

// NewService returns a new service instance
func NewService(repo repository.Interface, team teamservice.Interface) *Service {
	return &Service{repo: repo, team: team}
}

// Create creates a selections entry
func (s *Service) Create(selections model.Selections) error {
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
func (s *Service) List() ([]model.Selections, error) {
	repoSelections, err := s.repo.List()
	if err != nil {
		return nil, err
	}
	var selections []model.Selections
	for _, repoSelection := range repoSelections {
		var selection model.Selections
		for _, teamID := range repoSelection.TeamIDs {
			team, err := s.team.Get(teamID)
			if err != nil {
				return nil, err
			}
			selection.Teams = append(selection.Teams, team)
		}
		selections = append(selections, selection)
	}
	return selections, nil
}
