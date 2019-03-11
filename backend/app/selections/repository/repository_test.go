package repository

import (
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/suite"
	"github.com/tomney/finalfour/backend/app/selections"
	"github.com/tomney/finalfour/backend/app/team"
)

var selectionsStub = selections.Selections{
	Email: "testy@mctest.com",
	Teams: []team.Team{
		{ID: "test1", Name: "Test 1", ImageURL: "test1.png"},
		{ID: "test2", Name: "Test 2", ImageURL: "test2.png"},
		{ID: "test3", Name: "Test 3", ImageURL: "test3.png"},
		{ID: "test4", Name: "Test 4", ImageURL: "test4.png"},
	},
}

type createTestSuite struct {
	suite.Suite
	repo           *Repository
	mock           sqlmock.Sqlmock
	testSelections selections.Selections
}

func (s *createTestSuite) SetupTest() {
	db, mock, _ := sqlmock.New()
	s.repo = NewRepository(*db)
	s.mock = mock
	s.testSelections = selectionsStub
}

func (s *createTestSuite) TestCreateFailsIFPrepareStatementFails() {
	insertStatement := `
	INSERT INTO selections (email, first_pick, second_pick, third_pick, fourth_pick, created)
		VALUES(?, ?, ?, ?, ?, DATE(NOW()))
	`
	err := s.repo.Create(s.testSelections)
	expectedError := fmt.Errorf("some error")
	s.mock.ExpectPrepare(insertStatement).WillReturnError(expectedError)
	s.Assert().EqualError(err, expectedError.Error())
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(createTestSuite))
}
