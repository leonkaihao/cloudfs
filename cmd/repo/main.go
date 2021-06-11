package main

import (
	"os"

	"github.com/leonkaihao/cloudfs/cmd/repo/cmd"
	log "github.com/sirupsen/logrus"
)

func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)
}

func main() {
	if err := cmd.Execute(); err != nil {
		log.Errorln(err)
	}
}
