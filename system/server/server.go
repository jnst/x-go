package server

import (
	"fmt"
	"net/http"
)

// Server represents web server.
type Server struct {
	handler func(w http.ResponseWriter, r *http.Request)
}

// NewServer creates server instance.
func NewServer(handler func(w http.ResponseWriter, r *http.Request)) Server {
	return Server{handler: handler}
}

// Run runs server.
func (s Server) Run() {
	fmt.Println("[server] running..")
	http.HandleFunc("/", s.handler)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
