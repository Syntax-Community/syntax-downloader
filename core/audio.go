package core

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func ExtractMP3(url, output string) error {
	if output == "" {
		output = "downloads"
	}

	if err := os.MkdirAll(output, 0755); err != nil {
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

	outputTemplate := filepath.Join(output, "%(title)s.%(ext)s")

	cmd := exec.Command(
		"yt-dlp",
		"-x",
		"--quiet",
		"--no-warnings",
		"--progress",
		"--audio-format", "mp3",
		"--audio-quality", "0",
		"--user-agent", "Mozilla/5.0",
		"--no-part",
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
