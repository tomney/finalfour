package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/tomney/finalfour/backend/app/handler"
	selectionsapi "github.com/tomney/finalfour/backend/app/selections/api"
	selectionsrepo "github.com/tomney/finalfour/backend/app/selections/repository"
	selectionsservice "github.com/tomney/finalfour/backend/app/selections/service"
	teamservice "github.com/tomney/finalfour/backend/app/team/service"
)

//Tentative main
func main() {
	ctx := context.Background()
	log.Printf(": %v", ctx)

	port := "80"
	if p := os.Getenv("PORT"); p != "" {
		port = p
	}

	// Get the SQL database
	db := GetSQLDB()

	// Repository Layers
	selectionsRepo := selectionsrepo.NewRepository(db)

	// Service Layers
	teamService := teamservice.NewService()
	selectionsService := selectionsservice.NewService(selectionsRepo, teamService)

	// Handler Layers
	selectionsAPI := selectionsapi.NewHandler(selectionsService)

	r := mux.NewRouter()
	r.Methods("GET").Path("/api/v1/hello").Handler(handler.AppHandler(helloHandler))
	r.Methods("POST").Path("/api/v1/setSelection").Handler(handler.AppHandler(selectionsAPI.SubmitSelectionsHandler))
	r.Methods("GET").Path("/api/v1/listSelections").Handler(handler.AppHandler(selectionsAPI.ListSelectionsHandler))

	http.Handle("/", handlers.CombinedLoggingHandler(os.Stderr, r))

	log.Fatal(http.ListenAndServe(":"+port, nil))
}

// TODO this function is benign and should be removed after testing or replaced with a proper health endpoint
func helloHandler(w http.ResponseWriter, r *http.Request) *handler.AppError {
	w.Write([]byte("{\"greeting\":\"Bienvenidos\"}"))
	w.WriteHeader(200)
	return nil
}
