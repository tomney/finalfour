package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/tomney/finalfour/backend/app/handler"
	model "github.com/tomney/finalfour/backend/app/selections"
	"github.com/tomney/finalfour/backend/app/selections/service"
	"github.com/tomney/finalfour/backend/app/team"
)

// Interface implements the methods to interact with selections
type Interface interface {
	SubmitSelectionHandler(w http.ResponseWriter, r *http.Request) *handler.AppError
	ListSelectionsHandler(w http.ResponseWriter) *handler.AppError
}

// Handler handles requests for selections
type Handler struct {
	service service.Interface
}

// NewHandler returns a new service instance
func NewHandler(service service.Interface) *Handler {
	return &Handler{service: service}
}

// SubmitSelectionsHandler handles requests to submit final four selections
func (h *Handler) SubmitSelectionsHandler(w http.ResponseWriter, r *http.Request) *handler.AppError {
	log.Printf("Submit selection handler called")
	var selectionsRequest SubmitSelectionsRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&selectionsRequest)
	if err != nil {
		panic(err)
	}
	selections := model.Selections{
		Email: selectionsRequest.Email,
		Teams: []team.Team{
			{ID: selectionsRequest.Teams[0].ID, ImageURL: selectionsRequest.Teams[0].ImageURL, Name: selectionsRequest.Teams[0].Name},
			{ID: selectionsRequest.Teams[1].ID, ImageURL: selectionsRequest.Teams[1].ImageURL, Name: selectionsRequest.Teams[1].Name},
			{ID: selectionsRequest.Teams[2].ID, ImageURL: selectionsRequest.Teams[2].ImageURL, Name: selectionsRequest.Teams[2].Name},
			{ID: selectionsRequest.Teams[3].ID, ImageURL: selectionsRequest.Teams[3].ImageURL, Name: selectionsRequest.Teams[3].Name},
		},
	}

	err = h.service.Create(selections)
	if err != nil {
		log.Printf("An error occurred trying to create the selections entry.")
		return handler.AppErrorf(err, 500, err.Error())
	}

	w.Write([]byte("{}"))
	w.WriteHeader(200)
	return nil
}

// ListSelectionsHandler handles requests to submit final four selections
func (h *Handler) ListSelectionsHandler(w http.ResponseWriter, r *http.Request) *handler.AppError {
	log.Printf("List selections handler called")

	//Hardcode selection responses for now
	selection1 := model.Stub
	selection2 := model.Stub
	selections := []model.Selections{selection1, selection2}

	_, err := h.service.List()
	if err != nil {
		log.Printf("An error occurred trying to list the selections.")
		return handler.AppErrorf(err, 500, err.Error())
	}

	responseBody, err := json.Marshal(selections)
	if err != nil {
		log.Printf("Unable to marshal the selection list to JSON")
		return handler.AppErrorf(err, 500, "Unable to marshal the selection list to JSON: %v", err)
	}
	w.Write(responseBody)
	w.WriteHeader(200)
	return nil
}
