package authenticationservice

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func Route(mux *chi.Mux) {

	mux.Get("/", handler)

}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("helllo"))
}
