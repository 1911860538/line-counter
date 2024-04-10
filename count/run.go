package count

import (
	"log"
)

func Run() {
	params, err := parseFlags()
	handleErr(err)

	files, err := getFiles(params)
	handleErr(err)

	err = multiSetLineCount(files)
	handleErr(err)

	staticSlice := static(files)

	output(staticSlice)
}

func handleErr(err error) {
	if err != nil {
		log.Fatalf("ERROR: %v\n", err)
	}
}
