package selections

import (
	"log"
	"net/http"

	"github.com/tomney/angular-go-webapp/backend/app/handler"
)

// SubmitSelectionHandler handles requests to submit final four selections
func SubmitSelectionHandler(w http.ResponseWriter, r *http.Request) *handler.AppError {
	log.Printf("Submit selection handler called")
	w.Write([]byte("{}"))
	w.WriteHeader(200)
	return nil
}
