package main

import (
	"log"
	"net/http"
	"os"

	"github.com/angadsharma1016/c2c/model"
	"github.com/angadsharma1016/c2c/server/controller"
)

type server struct {
	host string
	port string
}

func (s server) serve(h *http.Handler) {
	log.Printf("Listening on %s%s", s.host, s.port)
	log.Fatal(http.ListenAndServe(s.host+s.port, *h))
}
func main() {
	con := model.Connect()
	defer con.Close()
	log.Println(os.Args)
	s := server{"0.0.0.0", ":" + os.Args[2]}
	s.serve(controller.Startup())
}
