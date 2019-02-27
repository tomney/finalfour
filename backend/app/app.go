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
)

//Tentative main
func main() {
	ctx := context.Background()
	log.Printf("%v", ctx)

	port := "80"
	if p := os.Getenv("PORT"); p != "" {
		port = p
	}

	// Repository Layers
	selectionsRepo := selectionsrepo.NewRepository("")

	// Service Layers
	selectionsService := selectionsservice.NewService(selectionsRepo)

	// Handler Layers
	selectionsAPI := selectionsapi.NewHandler(selectionsService)

	r := mux.NewRouter()
	r.Methods("GET").Path("/api/v1/hello").Handler(handler.AppHandler(helloHandler))
	r.Methods("POST").Path("/api/v1/setSelection").Handler(handler.AppHandler(selectionsAPI.SubmitSelectionsHandler))

	http.Handle("/", catchAll(r))

	log.Fatal(http.ListenAndServe(":"+port, nil))
}

// TODO this function is benign and should be removed after testing or replaced with a proper health endpoint
func helloHandler(w http.ResponseWriter, r *http.Request) *handler.AppError {
	w.Write([]byte("{\"greeting\":\"Bienvenidos\"}"))
	w.WriteHeader(200)
	return nil
}

func catchAll(r *mux.Router) http.Handler {
	log.Printf("Well this shit isnt working but at least you are catching results")
	return handlers.CombinedLoggingHandler(os.Stderr, r)
}
