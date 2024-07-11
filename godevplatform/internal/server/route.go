package server

import (
	"log"
	"net/http"
)

func init() {
	log.Println("route")
}

type Entrance struct{}

func NewEntrance() *Entrance {
	return &Entrance{}
}
func (e *Entrance) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println("path:", r.URL.Path)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(r.URL.Path))
}
