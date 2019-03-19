package service

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"github.com/tomney/finalfour/backend/app/selections"
	repomocks "github.com/tomney/finalfour/backend/app/selections/repository/mocks"
)

type createTestSuite struct {
	suite.Suite
	repo           *repomocks.Interface
	service        *Service
	selectionsStub selections.Selections
	teamIDs        []string
}

func (s *createTestSuite) SetupTest() {
	s.repo = &repomocks.Interface{}
	s.service = NewService(s.repo)
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
