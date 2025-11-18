package webhooks

import (
	"encoding/json"
)

// ArtifactHubEvent represents the top-level CloudEvents-style envelope sent by Artifact Hub.
type ArtifactHubEvent struct {
	SpecVersion     string          `json:"specversion"`
	ID              string          `json:"id"`
	Source          string          `json:"source"`
	Type            string          `json:"type"`
	DataContentType string          `json:"datacontenttype"`
	Data            ArtifactHubData `json:"data"`
}

// ArtifactHubData holds the payload data for the event.
type ArtifactHubData struct {
	Package ArtifactHubPackage `json:"package"`
}

// ArtifactHubPackage describes the package information inside the event.
type ArtifactHubPackage struct {
	Name                    string                `json:"name"`
	Version                 string                `json:"version"`
	URL                     string                `json:"url"`
	Changes                 []string              `json:"changes"`
	ContainsSecurityUpdates bool                  `json:"containsSecurityUpdates"`
	Prerelease              bool                  `json:"prerelease"`
	Repository              ArtifactHubRepository `json:"repository"`
}

// ArtifactHubRepository describes the repository metadata for the package.
type ArtifactHubRepository struct {
	Kind      string `json:"kind"`
	Name      string `json:"name"`
	Publisher string `json:"publisher"`
}

// ParseArtifactHubEvent unmarshals a JSON payload into ArtifactHubEvent.
func ParseArtifactHubEvent(b []byte) (*ArtifactHubEvent, error) {
	var ev ArtifactHubEvent
	if err := json.Unmarshal(b, &ev); err != nil {
		return nil, err
	}
	return &ev, nil
}
