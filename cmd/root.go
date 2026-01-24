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

	if os.Args[1] == "-h" || os.Args[1] == "--help" {
		showHelp()
		return
	}

	url := os.Args[1]
	flags := os.Args[2:]

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
		case "-h", "--help":
			showHelp()
			return
		}
	}

	output = utils.ExpandPath(output)
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
  -i           Show media info only
  -o <folder>  Output directory
  -h, --help   Show this help

Examples:
  get https://youtu.be/xyz
  get https://youtu.be/xyz -mp3
  get https://youtu.be/xyz -o ~/Downloads
  get https://youtu.be/xyz -i
`)
}
