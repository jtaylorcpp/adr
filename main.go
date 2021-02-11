package main

import (
	log "github.com/sirupsen/logrus"

	"github.com/jtaylorcpp/adr/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatalf("adr exited with and error: %s", err.Error())
	}
}
