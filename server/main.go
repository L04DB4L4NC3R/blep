package main

import (
	"log"
	"net/http"
)

type server struct {
	host string
	port string
}

func (s server) serve(h http.Handler) {
	log.Printf("Listening on %s%s", s.host, s.port)
	http.ListenAndServe(s.host+s.port, h)
}
func main() {
	s := server{"0.0.0.0", ":3000"}
	s.serve(nil)
}
