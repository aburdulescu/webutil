package webutil

import (
	"log"
	"net/http"
)

func ErrorHandler(f func(w http.ResponseWriter, r *http.Request) (int, error)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		statusCode, err := f(w, r)
		if err != nil {
			rspData := struct {
				Error string `json:"error"`
			}{
				Error: err.Error(),
			}
			Respond(w, r, statusCode, rspData)
			log.Printf("%q: %v", r.Method+" "+r.RequestURI, err)
		}
	}
}
