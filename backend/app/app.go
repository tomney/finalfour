package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
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
	r.Methods("GET").Path("/api/v1/hello").Handler(appHandler(helloHandler))
	http.Handle("/", catchAll(r))

	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func helloHandler(w http.ResponseWriter, r *http.Request) *appError {
	log.Printf("Endpoint was actually called. I know right!")
	w.Write([]byte("{\"greeting\":\"Bienvenidos\"}"))
	w.WriteHeader(200)
	return nil
}

func catchAll(r *mux.Router) http.Handler {
	log.Printf("Well this shit isnt working but at least you are catching results")
	return handlers.CombinedLoggingHandler(os.Stderr, r)
}

type appHandler func(http.ResponseWriter, *http.Request) *appError

type appError struct {
	Error   error
	Message string
	Code    int
}

func (fn appHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if e := fn(w, r); e != nil { // e is *appError, not os.Error.
		log.Printf("Handler error: status code: %d, message: %s, underlying err: %#v",
			e.Code, e.Message, e.Error)

		http.Error(w, e.Message, e.Code)
	}
}

func appErrorf(err error, format string, v ...interface{}) *appError {
	return &appError{
		Error:   err,
		Message: fmt.Sprintf(format, v...),
		Code:    500,
	}
}
