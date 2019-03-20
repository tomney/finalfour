package service

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"github.com/tomney/finalfour/backend/app/selections"
	"github.com/tomney/finalfour/backend/app/selections/repository"
	repomocks "github.com/tomney/finalfour/backend/app/selections/repository/mocks"
	teamservice "github.com/tomney/finalfour/backend/app/team/service/mocks"
)

type createTestSuite struct {
	suite.Suite
	repo           *repomocks.Interface
	team           *teamservice.Interface
	service        *Service
	selectionsStub selections.Selections
	teamIDs        []string
}

func (s *createTestSuite) SetupTest() {
	s.repo = &repomocks.Interface{}
	s.team = &teamservice.Interface{}
	s.service = NewService(s.repo, s.team)
	s.selectionsStub = selections.Stub
	s.teamIDs = []string{
		s.selectionsStub.Teams[0].ID,
		s.selectionsStub.Teams[1].ID,
		s.selectionsStub.Teams[2].ID,
		s.selectionsStub.Teams[3].ID,
	}
}

func (s *createTestSuite) TestReturnsErrorIfCallToGetSelectionFails() {
	expectedError := fmt.Errorf("Can't get it")
	s.repo.On("Get", mock.Anything).Return(nil, expectedError)
	err := s.service.Create(s.selectionsStub)
	s.Assert().EqualError(err, expectedError.Error())
}

func (s *createTestSuite) TestDeletesExistingSelectionIfSelectionAlreadyExistsForEmail() {
	s.repo.On("Get", mock.Anything).Return(s.teamIDs, nil)
	s.repo.On("Delete", mock.Anything).Return(nil)
	s.repo.On("Create", mock.Anything).Return(nil)
	s.service.Create(s.selectionsStub)
	s.repo.AssertCalled(s.T(), "Delete", s.selectionsStub.Email)
}

func (s *createTestSuite) TestReturnsErrorIfUnableToDeleteExistingEntry() {
	expectedError := fmt.Errorf("Can't delete it")
	s.repo.On("Get", mock.Anything).Return(s.teamIDs, nil)
	s.repo.On("Delete", mock.Anything).Return(expectedError)
	err := s.service.Create(s.selectionsStub)
	s.Assert().EqualError(err, expectedError.Error())
}

func TestCreateTestSuite(t *testing.T) {
	suite.Run(t, new(createTestSuite))
}

type listTestSuite struct {
	suite.Suite
	repo           *repomocks.Interface
	team           *teamservice.Interface
	service        *Service
	selectionsStub selections.Selections
	teamIDs        []string
}

func (s *listTestSuite) SetupTest() {
	s.repo = &repomocks.Interface{}
	s.team = &teamservice.Interface{}
	s.service = NewService(s.repo, s.team)
	s.selectionsStub = selections.Stub
	s.teamIDs = []string{
		s.selectionsStub.Teams[0].ID,
		s.selectionsStub.Teams[1].ID,
		s.selectionsStub.Teams[2].ID,
		s.selectionsStub.Teams[3].ID,
	}
}

func (s *listTestSuite) TestReturnsErrorIfCallToListFails() {
	expectedError := fmt.Errorf("Can't list it")
	s.repo.On("List", mock.Anything).Return(nil, expectedError)
	selections, err := s.service.List()
	s.Assert().Nil(selections)
	s.Assert().EqualError(err, expectedError.Error())
}

func (s *listTestSuite) TestReturnsErrorIfCallToGetTeamFails() {
	expectedError := fmt.Errorf("Can't get it")
	s.repo.On("List", mock.Anything).
		Return([]repository.Selections{repository.SelectionsStub}, expectedError)
	s.team.On("Get", mock.Anything).Return(nil, expectedError)
	selections, err := s.service.List()
	s.Assert().Nil(selections)
	s.Assert().EqualError(err, expectedError.Error())
}

func (s *listTestSuite) TestReturnsSelections() {
	repositorySelections := []repository.Selections{
		{
			Email:   selections.Stub.Email,
			TeamIDs: []string{selections.Stub.Teams[0].ID, selections.Stub.Teams[1].ID, selections.Stub.Teams[2].ID, selections.Stub.Teams[3].ID},
			Created: "",
		},
	}

	s.repo.On("List", mock.Anything).
		Return(repositorySelections, nil)
	s.team.On("Get", selections.Stub.Teams[0].ID).Return(selections.Stub.Teams[0], nil)
	s.team.On("Get", selections.Stub.Teams[1].ID).Return(selections.Stub.Teams[1], nil)
	s.team.On("Get", selections.Stub.Teams[2].ID).Return(selections.Stub.Teams[2], nil)
	s.team.On("Get", selections.Stub.Teams[3].ID).Return(selections.Stub.Teams[3], nil)

	allSelections, err := s.service.List()
	s.Assert().EqualValues(allSelections, []selections.Selections{selections.Stub})
	s.Assert().Nil(err)
}
func TestListTestSuite(t *testing.T) {
	suite.Run(t, new(listTestSuite))
}
