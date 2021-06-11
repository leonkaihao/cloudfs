package cmd

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var parseCmd = &cobra.Command{
	Use:   "parse",
	Short: "Parse a file/folder to the repo",
	Long:  "Parse a file/folder and record all the file(s) status to the repo",
	Run: func(cmd *cobra.Command, args []string) {
		log.Error("need implementation")
	},
}
