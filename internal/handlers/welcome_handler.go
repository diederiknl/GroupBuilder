package handlers

import (
	"net/http"
)

func Welcome(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Welcome to the GroupBuilder API!"))
}
