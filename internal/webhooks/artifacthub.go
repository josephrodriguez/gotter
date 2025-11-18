package webhooks

import (
	"io"
	"log"
	"net/http"
	"strings"
)

func ArtifactHubHandler(w http.ResponseWriter, r *http.Request) {
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

	ev, err := ParseArtifactHubEvent(body)
	if err != nil {
		log.Printf("[Gotter] failed to parse Artifact Hub event: %v", err)
		http.Error(w, "Invalid payload", http.StatusBadRequest)
		return
	}

	pkg := ev.Data.Package
	changes := strings.Join(pkg.Changes, "; ")

	log.Printf("[Gotter] ArtifactHub event received: id=%s type=%s source=%s spec=%s", ev.ID, ev.Type, ev.Source, ev.SpecVersion)
	log.Printf("[Gotter] Package: name=%s version=%s url=%s repo=%s/%s publisher=%s prerelease=%t securityUpdates=%t changes=%s",
		pkg.Name, pkg.Version, pkg.URL, pkg.Repository.Kind, pkg.Repository.Name, pkg.Repository.Publisher, pkg.Prerelease, pkg.ContainsSecurityUpdates, changes)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ArtifactHub webhook received"))
}

// (types and parsing live in artifacthub_types.go)
