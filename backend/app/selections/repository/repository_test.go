package repository

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/suite"
	"github.com/tomney/finalfour/backend/app/selections"
)

type createTestSuite struct {
	suite.Suite
	repo           Interface
	mock           sqlmock.Sqlmock
	testSelections selections.Selections
}

func (s *createTestSuite) SetupTest() {
	db, mock, _ := sqlmock.New()
	s.repo = NewRepository(*db)
	s.mock = mock
	s.testSelections = selections.Stub
}

func (s *createTestSuite) TestCreateFailsIfPrepareStatementFails() {
	insertStatement := "INSERT INTO selections"
	expectedError := fmt.Errorf("some preparation error")
	s.mock.ExpectPrepare(insertStatement).WillReturnError(expectedError)

	err := s.repo.Create(s.testSelections)
	s.Assert().EqualError(err, expectedError.Error())
}

func (s *createTestSuite) TestCreateFailsIfStatementExecutionFails() {
	insertStatement := "INSERT INTO selections"
	expectedError := fmt.Errorf("some execution error")
	s.mock.ExpectPrepare(insertStatement)
	s.mock.ExpectPrepare(insertStatement).
		ExpectExec().
		WithArgs(s.testSelections.Email, s.testSelections.Teams[0].ID, s.testSelections.Teams[1].ID, s.testSelections.Teams[2].ID, s.testSelections.Teams[3].ID).
		WillReturnError(expectedError)
	err := s.repo.Create(s.testSelections)
	s.Assert().EqualError(err, expectedError.Error())
}

func (s *createTestSuite) TestCreateSucceeds() {
	insertStatement := "INSERT INTO selections"
	s.mock.ExpectPrepare(insertStatement)
	s.mock.ExpectPrepare(insertStatement).
		ExpectExec().
		WithArgs(s.testSelections.Email, s.testSelections.Teams[0].ID, s.testSelections.Teams[1].ID, s.testSelections.Teams[2].ID, s.testSelections.Teams[3].ID).
		WillReturnResult(sqlmock.NewResult(0, 1))
	err := s.repo.Create(s.testSelections)
	s.Assert().Nil(err)
}

func TestCreateTestSuite(t *testing.T) {
	suite.Run(t, new(createTestSuite))
}

type deleteTestSuite struct {
	suite.Suite
	repo           Interface
	mock           sqlmock.Sqlmock
	testSelections selections.Selections
}

func (s *deleteTestSuite) SetupTest() {
	db, mock, _ := sqlmock.New()
	s.repo = NewRepository(*db)
	s.mock = mock
	s.testSelections = selections.Stub
}

func (s *deleteTestSuite) TestDeleteFailsIfPrepareStatementFails() {
	deleteStatement := "DELETE FROM selections"
	expectedError := fmt.Errorf("some preparation error")
	s.mock.ExpectPrepare(deleteStatement).WillReturnError(expectedError)

	err := s.repo.Delete(s.testSelections.Email)
	s.Assert().EqualError(err, expectedError.Error())
}

func (s *deleteTestSuite) TestCreateFailsIfStatementExecutionFails() {
	deleteStatement := "DELETE FROM selections"
	expectedError := fmt.Errorf("some execution error")
	s.mock.ExpectPrepare(deleteStatement)
	s.mock.ExpectPrepare(deleteStatement).
		ExpectExec().
		WithArgs(s.testSelections.Email).
		WillReturnError(expectedError)
	err := s.repo.Delete(s.testSelections.Email)
	s.Assert().EqualError(err, expectedError.Error())
}

func (s *deleteTestSuite) TestDeleteSucceeds() {
	deleteStatement := "DELETE FROM selections"
	s.mock.ExpectPrepare(deleteStatement)
	s.mock.ExpectPrepare(deleteStatement).
		ExpectExec().
		WithArgs(s.testSelections.Email).
		WillReturnResult(sqlmock.NewResult(0, 1))
	err := s.repo.Delete(s.testSelections.Email)
	s.Assert().Nil(err)
}

func TestDeleteTestSuite(t *testing.T) {
	suite.Run(t, new(deleteTestSuite))
}

type getTestSuite struct {
	suite.Suite
	repo           Interface
	mock           sqlmock.Sqlmock
	testSelections selections.Selections
}

func (s *getTestSuite) SetupTest() {
	db, mock, _ := sqlmock.New()
	s.repo = NewRepository(*db)
	s.mock = mock
	s.testSelections = selections.Stub
}

func (s *getTestSuite) TestGetReturnsNilIfQueryComesBackEmpty() {
	selectStatement := "SELECT .* FROM selections"
	s.mock.ExpectQuery(selectStatement).WillReturnError(sql.ErrNoRows)

	teamIDs, err := s.repo.Get(s.testSelections.Email)
	s.Assert().Nil(teamIDs)
	s.Assert().Nil(err)
}

func (s *getTestSuite) TestGetReturnsErrorIfQueryErrors() {
	selectStatement := "SELECT .* FROM selections"
	expectedError := fmt.Errorf("shits broke")
	s.mock.ExpectQuery(selectStatement).WillReturnError(expectedError)

	teamIDs, err := s.repo.Get(s.testSelections.Email)
	s.Assert().Nil(teamIDs)
	s.Assert().EqualError(err, expectedError.Error())
}

func (s *getTestSuite) TestGetReturnsResponseIfQueryIsSuccessful() {
	selectStatement := "SELECT .* FROM selections"
	columns := []string{"email", "first", "second", "third", "fourth", "created"}
	expectedIDs := []string{s.testSelections.Teams[0].ID, s.testSelections.Teams[1].ID, s.testSelections.Teams[2].ID, s.testSelections.Teams[3].ID}
	rows := sqlmock.NewRows(columns).AddRow(s.testSelections.Email, s.testSelections.Teams[0].ID, s.testSelections.Teams[1].ID, s.testSelections.Teams[2].ID, s.testSelections.Teams[3].ID, "2019-01-01")
	s.mock.ExpectQuery(selectStatement).WithArgs(s.testSelections.Email).WillReturnRows(rows)
	teamIDs, err := s.repo.Get(s.testSelections.Email)
	s.Assert().EqualValues(teamIDs, expectedIDs)
	s.Assert().Nil(err)
}

func TestGetTestSuite(t *testing.T) {
	suite.Run(t, new(getTestSuite))
}

type listTestSuite struct {
	suite.Suite
	repo           Interface
	mock           sqlmock.Sqlmock
	testSelections selections.Selections
}

func (s *listTestSuite) SetupTest() {
	db, mock, _ := sqlmock.New()
	s.repo = NewRepository(*db)
	s.mock = mock
	s.testSelections = selections.Stub
}

func (s *listTestSuite) TestReturnsNilIfQueryComesBackEmpty() {
	selectStatement := "SELECT .* FROM selections"
	s.mock.ExpectQuery(selectStatement).WillReturnError(sql.ErrNoRows)

	allSelections, err := s.repo.List()
	s.Assert().Nil(allSelections)
	s.Assert().Nil(err)
}

func (s *listTestSuite) TestReturnsErrIfQueryReturnsError() {
	selectStatement := "SELECT .* FROM selections"
	expectedError := fmt.Errorf("Can't list it")
	s.mock.ExpectQuery(selectStatement).WillReturnError(expectedError)

	allSelections, err := s.repo.List()
	s.Assert().Nil(allSelections)
	s.Assert().EqualError(err, expectedError.Error())
}

func (s *listTestSuite) TestAppendsSelections() {
	selectStatement := "SELECT .* FROM selections"
	columns := []string{"email", "first", "second", "third", "fourth", "created"}
	rows := sqlmock.NewRows(columns).
		AddRow(SelectionsStub.Email, SelectionsStub.TeamIDs[0], SelectionsStub.TeamIDs[1], SelectionsStub.TeamIDs[2], SelectionsStub.TeamIDs[3], SelectionsStub.Created).
		AddRow(SelectionsStub.Email, SelectionsStub.TeamIDs[0], SelectionsStub.TeamIDs[1], SelectionsStub.TeamIDs[2], SelectionsStub.TeamIDs[3], SelectionsStub.Created)
	expected := []Selections{SelectionsStub, SelectionsStub}
	s.mock.ExpectQuery(selectStatement).WillReturnRows(rows)
	allSelections, err := s.repo.List()
	s.Assert().Nil(err)
	s.Assert().EqualValues(allSelections, expected)
}

func TestListTestSuite(t *testing.T) {
	suite.Run(t, new(listTestSuite))
}
