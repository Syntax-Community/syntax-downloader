package core

import (
	"encoding/json"
	"fmt"
	"os/exec"
)

type VideoInfo struct {
	Title      string `json:"title"`
	Uploader   string `json:"uploader"`
	Duration   int    `json:"duration"`
}

func Info(url string) error {
	cmd := exec.Command(
		"yt-dlp",
		"--quiet",
		"--no-warnings",
		"--dump-single-json",
		url,
	)

	out, err := cmd.Output()
	if err != nil {
		return err
	}

	var info VideoInfo
	if err := json.Unmarshal(out, &info); err != nil {
		return err
	}

	fmt.Println("Title   :", info.Title)
	fmt.Println("Uploader:", info.Uploader)
	fmt.Println("Duration:", formatDuration(info.Duration))

	return nil
}
func formatDuration(sec int) string {
	m := sec / 60
	s := sec % 60
	return fmt.Sprintf("%02d:%02d", m, s)
}

func formatDate(d string) string {
	if len(d) != 8 {
		return d
	}
	return fmt.Sprintf("%s-%s-%s", d[:4], d[4:6], d[6:])
}
