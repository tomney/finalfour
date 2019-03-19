package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"net/http/httptest"

	"github.com/stretchr/testify/suite"
	"github.com/tomney/finalfour/backend/app/handler"
	"github.com/tomney/finalfour/backend/app/selections"
	servicemocks "github.com/tomney/finalfour/backend/app/selections/service/mocks"
)

type createTestSuite struct {
	suite.Suite
	service            *servicemocks.Interface
	handler            *Handler
	selectionsStub     selections.Selections
	httpResponseWriter http.ResponseWriter
	requestBody        SubmitSelectionsRequest
}

func (s *createTestSuite) SetupTest() {
	s.service = &servicemocks.Interface{}
	s.handler = NewHandler(s.service)
	s.selectionsStub = selections.Stub
	s.httpResponseWriter = &httptest.ResponseRecorder{}
	s.requestBody = SubmitSelectionsRequest{
		Email: selections.Stub.Email,
		Teams: []TeamRequest{
			{ID: selections.Stub.Teams[0].ID, ImageURL: selections.Stub.Teams[0].ImageURL, Name: selections.Stub.Teams[0].Name},
			{ID: selections.Stub.Teams[1].ID, ImageURL: selections.Stub.Teams[1].ImageURL, Name: selections.Stub.Teams[1].Name},
			{ID: selections.Stub.Teams[2].ID, ImageURL: selections.Stub.Teams[2].ImageURL, Name: selections.Stub.Teams[2].Name},
			{ID: selections.Stub.Teams[3].ID, ImageURL: selections.Stub.Teams[3].ImageURL, Name: selections.Stub.Teams[3].Name},
		},
	}
}

func (s *createTestSuite) TestReturnsErrorIfServiceReturnsError() {
	jsonBody, _ := json.Marshal(s.requestBody)
	request := httptest.NewRequest("POST", "/api/v1/setSelection", bytes.NewReader(jsonBody))
	expectedError := fmt.Errorf("Can't create it")
	expectedHandlerError := handler.AppErrorf(expectedError, 500, expectedError.Error())
	s.service.On("Create", s.selectionsStub).Return(expectedError)
	handlerError := s.handler.SubmitSelectionsHandler(s.httpResponseWriter, request)
	s.Assert().Equal(handlerError, expectedHandlerError)
}

func TestCreateTestSuite(t *testing.T) {
	suite.Run(t, new(createTestSuite))
}
