package parse

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var FileCmd = &cobra.Command{
	Use:   "file",
	Short: "Parse a file to the repo",
	Long:  "Parse and record the file status to the repo",
	Run: func(cmd *cobra.Command, args []string) {
		log.Error("need implementation")
	},
}
