package controller

import (
	"encoding/json"
	"net/http"

	"github.com/rs/cors"
)

func Startup() *http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/v1", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(struct {
			Status bool
		}{true})
	})
	mux.HandleFunc("/api/v1/read", readLogs())
	corsMux := cors.Default().Handler(mux)
	return &corsMux
}
