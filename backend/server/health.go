package server

import (
	"net/http"
)

func health(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusNoContent)
}

func Start() {
	http.HandleFunc("/health", health)
	http.ListenAndServe(":8080", nil)
}
