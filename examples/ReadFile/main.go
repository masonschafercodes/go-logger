package main

import (
	"os"

	log "github.com/masonschafercodes/go-logger"
)

func main() {

	dat, err := os.ReadFile("./fake_data.txt")

	if err != nil {
		log.Error(err)
	}

	log.Success(string(dat))

}
