package webutil

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func Respond(w http.ResponseWriter, r *http.Request, status int, data interface{}) {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if _, err := io.Copy(w, &buf); err != nil {
		log.Println("respond:", err)
	}
}
