package utils

import (
	"os"
	"path/filepath"
	"strings"
)

func ExpandPath(path string) string {
	if strings.HasPrefix(path, "~/") {
		home, err := os.UserHomeDir()
		if err == nil {
			return filepath.Join(home, path[2:])
		}
	}
	return path
}

func DefaultDownloadDir() string {
	home, err := os.UserHomeDir()
	if err != nil || home == "" {
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

	return home
}
