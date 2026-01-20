package core

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func Download(url, outputDir string) error {
	if outputDir == "" {
		outputDir = "downloads"
	}

	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return err
	}

	titleCmd := exec.Command(
		"yt-dlp",
		"--no-warnings",
		"--print", "title",
		"--skip-download",
		url,
	)

	titleOut, err := titleCmd.Output()
	if err != nil {
		return err
	}

	title := strings.TrimSpace(string(titleOut))
	fmt.Println(title)

	outputTemplate := filepath.Join(outputDir, "%(title)s.%(ext)s")

	cmd := exec.Command(
		"yt-dlp",
		"--quiet",
		"--no-warnings",
		"--progress",
		"--merge-output-format", "mp4",
		"-f", "bestvideo+bestaudio/best",
		"-o", outputTemplate,
		url,
	)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stdout
	cmd.Stdin = os.Stdin

	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}
