package utils

import (
	"os"
	"path/filepath"
)

func DefaultDownloadDir() string {
	home := os.Getenv("HOME")
	if home == "" {
		home = "."
	}

	linuxDir := filepath.Join(home, "Downloads")
	if stat, err := os.Stat(linuxDir); err == nil && stat.IsDir() {
		return linuxDir
	}

	androidDir := filepath.Join(home, "storage", "downloads")
	if stat, err := os.Stat(androidDir); err == nil && stat.IsDir() {
		return androidDir
	}

	return "."
}
