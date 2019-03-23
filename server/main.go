package main

import (
	"log"
	"net/http"

	"github.com/angadsharma1016/c2c/model"
	"github.com/angadsharma1016/c2c/server/controller"
)

type server struct {
	host string
	port string
}

func (s server) serve(h *http.Handler) {
	log.Printf("Listening on %s%s", s.host, s.port)
	http.ListenAndServe(s.host+s.port, *h)
}
func main() {
	con := model.Connect()
	defer con.Close()
	s := server{"0.0.0.0", ":3000"}
	s.serve(controller.Startup())
}
