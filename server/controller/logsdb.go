package controller

import (
	"encoding/json"
	"net/http"

	"github.com/angadsharma1016/c2c/model"
)

func readLogs() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		c := make(chan []model.Logs)
		go model.Read(c)
		msg := <-c
		json.NewEncoder(w).Encode(msg)
	}
}
