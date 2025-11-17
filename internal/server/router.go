package server

import (
	"net/http"

	"github.com/josephrodriguez/gotter/internal/webhooks"
)

func NewRouter() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/webhooks/artifactory", webhooks.ArtifactoryHandler)
	return mux
}
