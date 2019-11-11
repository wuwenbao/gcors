package gocors

import (
	"log"
	"net/http"
	"testing"
)

func TestNew(t *testing.T) {

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("cors"))
	})

	log.Fatal(http.ListenAndServe(":8080", New(mux)))
}