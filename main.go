package main

import (
	"log"

	"github.com/josephrodriguez/gotter/internal/config"
	"github.com/josephrodriguez/gotter/internal/server"
)

func main() {
	host := config.GetEnvStr("GOTTER_HOST", "0.0.0.0")
	port := config.GetEnvInt("GOTTER_PORT", 8080)

	log.Println("[Gotter] Starting service...")

	if err := server.Start(host, port); err != nil {
		log.Fatalf("[Gotter] Server failed: %v", err)
	}
}
