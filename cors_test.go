package gcors

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

	//cors := New(mux) //默认参数为 *
	cors := New(
		mux,
		WithOrigin("*"),
		WithMethods("*"),
		WithMethods("*"),
	)

	log.Fatal(http.ListenAndServe(":8080", cors))
}
