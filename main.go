package main

import (
	"os"

	"github.com/elastic/beats/libbeat/beat"

	"github.com/bwarminski/kinesisbeat/beater"
)

func main() {
	err := beat.Run("kinesisbeat", "", beater.New)
	if err != nil {
		os.Exit(1)
	}
}
