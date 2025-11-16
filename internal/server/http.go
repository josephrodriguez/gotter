package server

import (
	"fmt"
	"net/http"
)

func Start(host string, port int) error {
	router := NewRouter()

	addr := fmt.Sprintf("%s:%d", host, port)
	fmt.Printf("[Gotter] Server running on http://%s\n", addr)

	return http.ListenAndServe(addr, router)
}
