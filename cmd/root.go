package cmd

import (
	"os"
	"get/core"
	"get/utils"
)

func Execute() {
	if len(os.Args) < 2 {
		utils.Error("get -h")
		return
	}

	for _, f := range os.Args[1:] {
		if f == "-h" || f == "--help" {
			showHelp()
			return
		}
	}

	url := ""
	flags := []string{}

	for i := 1; i < len(os.Args); i++ {
		arg := os.Args[i]
		if arg == "-mp3" || arg == "-i" || arg == "-o" {
			flags = append(flags, arg)
			if arg == "-o" && i+1 < len(os.Args) {
				flags = append(flags, os.Args[i+1])
				i++
			}
		} else if url == "" && arg[0] != '-' {
			url = arg
		}
	}

	if url == "" {
		utils.Error("get -h")
		return
	}

	output := utils.DefaultDownloadDir()
	mp3 := false
	info := false

	for i := 0; i < len(flags); i++ {
		switch flags[i] {
		case "-mp3":
			mp3 = true
		case "-i":
			info = true
		case "-o":
			if i+1 < len(flags) {
				output = flags[i+1]
				i++
			}
		}
	}

	utils.EnsureDir(output)

	if info {
		core.Info(url)
		return
	}

	if mp3 {
		if err := core.ExtractMP3(url, output); err != nil {
			utils.Error(err.Error())
		}
	} else {
		if err := core.Download(url, output); err != nil {
			utils.Error(err.Error())
		}
	}
}

func showHelp() {
	println(`
Usage: get <url> [options]

Options:
  -mp3         Extract audio as MP3
  -i           Show video/audio info only
  -o <folder>  Specify output folder
  -h, --help   Show this help message

Examples:
  get https://youtu.be/xyz123
  get https://youtu.be/xyz123 -mp3
  get https://youtu.be/xyz123 -mp3 -o ~/Downloads/Music
  get https://youtu.be/xyz123 -i
`)
}
