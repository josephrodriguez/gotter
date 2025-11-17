package server

import (
	"fmt"
	"log"
	"net/http"
)

func Start(host string, port int) error {
	router := NewRouter()

	addr := fmt.Sprintf("%s:%d", host, port)
	log.Printf("[Gotter] Server running on http://%s\n", addr)

	return http.ListenAndServe(addr, router)
}
