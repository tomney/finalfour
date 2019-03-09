package service

import (
	"github.com/tomney/finalfour/backend/app/team"
)

// Interface implements the methods to interact with selections
type Interface interface {
	Get(string) (team.Team, error)
}

// Service handles the collection and alteration of selections
type Service struct {
}

// NewService returns a new service instance
func NewService() *Service {
	return &Service{}
}

// Get retrieves the team for a given team ID
func (s *Service) Get(id string) (team.Team, error) {
	//TODO  build this function
	return team.Team{
		ID:       "gregsteam",
		ImageURL: "mybutt.png",
		Name:     "Greg's Team",
	}, nil
}
