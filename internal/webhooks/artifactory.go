package webhooks

import (
	"fmt"
	"io"
	"net/http"
)

func ArtifactoryHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST allowed", http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	fmt.Println("[Gotter] Received Artifactory webhook request:")
	fmt.Println(string(body))

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Artifactory webhook received"))
}
