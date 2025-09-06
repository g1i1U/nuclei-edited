package utils

import (
	"strings"

	"github.com/g1i1u/nuclei-edited/v3/pkg/catalog/config"
	"github.com/g1i1u/nuclei-edited/v3/pkg/keys"
)

const (
	// TemplatesRepoURL is the URL for files in nuclei-templates repository
	TemplatesRepoURL = "https://cloud.projectdiscovery.io/public/"
)

// TemplatePathURL returns the Path and URL for the provided template
func TemplatePathURL(fullPath, templateId, templateVerifier string) (path string, url string) {
	configData := config.DefaultConfig
	if configData.TemplatesDirectory != "" && strings.HasPrefix(fullPath, configData.TemplatesDirectory) {
		path = strings.TrimPrefix(strings.TrimPrefix(fullPath, configData.TemplatesDirectory), "/")
	}
	if templateVerifier == keys.PDVerifier {
		url = TemplatesRepoURL + templateId
	}
	return
}
