package handler

import (
	"fmt"
	"log"
	"net/http"
)

// AppHandler handles calls to the api
type AppHandler func(http.ResponseWriter, *http.Request) *AppError

// AppError is the struct for any error committed by the app
type AppError struct {
	Error   error
	Message string
	Code    int
}

// ServeHTTP allows an app handler to serve files over http
func (fn AppHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if e := fn(w, r); e != nil { // e is *appError, not os.Error.
		log.Printf("Handler error: status code: %d, message: %s, underlying err: %#v",
			e.Code, e.Message, e.Error)

		http.Error(w, e.Message, e.Code)
	}
}

//AppErrorf creates an app error
func AppErrorf(err error, code int, format string, v ...interface{}) *AppError {
	return &AppError{
		Error:   err,
		Message: fmt.Sprintf(format, v...),
		Code:    code,
	}
}
