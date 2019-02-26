package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/tomney/angular-go-webapp/backend/app/handler"
	"github.com/tomney/angular-go-webapp/backend/app/selections"
)

//Tentative main
func main() {
	ctx := context.Background()
	log.Printf("%v", ctx)

	port := "80"
	if p := os.Getenv("PORT"); p != "" {
		port = p
	}

	r := mux.NewRouter()
	r.Methods("GET").Path("/api/v1/hello").Handler(handler.AppHandler(helloHandler))
	r.Methods("POST").Path("/api/v1/setSelection").Handler(handler.AppHandler(selections.SubmitSelectionHandler))
	r.Methods("GET").Path("/api/v1/setSelection").Handler(handler.AppHandler(selections.SubmitSelectionHandler))

	http.Handle("/", catchAll(r))

	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func helloHandler(w http.ResponseWriter, r *http.Request) *handler.AppError {
	w.Write([]byte("{\"greeting\":\"Bienvenidos\"}"))
	w.WriteHeader(200)
	return nil
}

func catchAll(r *mux.Router) http.Handler {
	log.Printf("Well this shit isnt working but at least you are catching results")
	return handlers.CombinedLoggingHandler(os.Stderr, r)
}
