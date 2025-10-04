package version

import (
	_ "embed"
	"log"
	"strings"
)

//go:embed VERSION
var Version string

func init() {
	Version = strings.TrimSpace(Version)
	if Version == "" {
		log.Fatal("[Version] version is empty")
	}
	log.Printf("[Version] %s\n", Version)
}
