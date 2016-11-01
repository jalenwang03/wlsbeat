package main

import (
	"os"

	"github.com/elastic/beats/libbeat/beat"

	"github.com/jalenwang03/wlsbeat/beater"
)

func main() {
	err := beat.Run("wlsbeat", "", beater.New)
	if err != nil {
		os.Exit(1)
	}
}
