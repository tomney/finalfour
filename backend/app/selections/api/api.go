package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/tomney/finalfour/backend/app/handler"
	model "github.com/tomney/finalfour/backend/app/selections"
	"github.com/tomney/finalfour/backend/app/selections/service"
)

// Interface implements the methods to interact with selections
type Interface interface {
	SubmitSelectionHandler(model.Selections) error
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
		Teams: []model.Team{
			{ID: selectionsRequest.Teams[0].ID, ImageURL: selectionsRequest.Teams[0].ImageURL, Name: selectionsRequest.Teams[0].Name},
			{ID: selectionsRequest.Teams[1].ID, ImageURL: selectionsRequest.Teams[1].ImageURL, Name: selectionsRequest.Teams[1].Name},
			{ID: selectionsRequest.Teams[2].ID, ImageURL: selectionsRequest.Teams[2].ImageURL, Name: selectionsRequest.Teams[2].Name},
			{ID: selectionsRequest.Teams[3].ID, ImageURL: selectionsRequest.Teams[3].ImageURL, Name: selectionsRequest.Teams[3].Name},
		},
	}

	err = h.service.Create(selections)
	if err != nil {
		log.Printf("An error occurred trying to create the selections entry.")
	}

	w.Write([]byte("{}"))
	w.WriteHeader(200)
	return nil
}
