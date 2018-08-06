package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/apex/gateway"
)

// ContentType is the Content-Type header set in responses.
const ContentType = "application/json; charset=utf8"

// Message contains a simple message response.
type Message struct {
	Message string `json:"message"`
}

// Messages used by http.HandlerFunc functions.
var (
	Resp = Message{"A succesful connection has been made!"}
)

// ConnectionHandler is a http.HandlerFunc for the / path.
func ConnectionHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(Resp)
}

// h wraps a http.HandlerFunc and adds common headers.
func h(next http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", ContentType)
		next.ServeHTTP(w, r)
	})
}

func main() {
	http.Handle("/", h(ConnectionHandler))
	log.Fatal(gateway.ListenAndServe(":9000", nil))
}
