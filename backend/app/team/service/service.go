package service

import (
	"fmt"

	"github.com/tomney/finalfour/backend/app/team"
)

var teams = map[string]team.Team{
	"duke":          {ID: "duke", Name: "Duke", ImageURL: "duke.png"},
	"tennessee":     {ID: "tennessee", Name: "Tennessee", ImageURL: "tennessee.png"},
	"virginia":      {ID: "virginia", Name: "Virginia", ImageURL: "virginia.png"},
	"gonzaga":       {ID: "gonzaga", Name: "Gonzaga", ImageURL: "gonzaga.png"},
	"kentucky":      {ID: "kentucky", Name: "Kentucky", ImageURL: "kentucky.png"},
	"michigan":      {ID: "michigan", Name: "Michigan", ImageURL: "michigan.png"},
	"northcarolina": {ID: "northcarolina", Name: "North Carolina", ImageURL: "northcarolina.png"},
	"michiganstate": {ID: "michiganstate", Name: "Michigan State", ImageURL: "michiganstate.png"},
	"purdue":        {ID: "purdue", Name: "Purdue", ImageURL: "purdue.png"},
	"kansas":        {ID: "kansas", Name: "Kansas", ImageURL: "kansas.png"},
	"houston":       {ID: "houston", Name: "Houston", ImageURL: "houston.png"},
	"marquette":     {ID: "marquette", Name: "Marquette", ImageURL: "marquette.png"},
	"iowast":        {ID: "iowast", Name: "Iowa St", ImageURL: "iowast.png"},
	"nevada":        {ID: "nevada", Name: "Nevada", ImageURL: "nevada.png"},
	"louisville":    {ID: "louisville", Name: "Louisville", ImageURL: "louisville.png"},
	"wisconsin":     {ID: "wisconsin", Name: "Wisconsin", ImageURL: "wisconsin.png"},
}

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
	if _, ok := teams[id]; !ok {
		return team.Team{}, fmt.Errorf("team does not exist")
	}

	return teams[id], nil
}
