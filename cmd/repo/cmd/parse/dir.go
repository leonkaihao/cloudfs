package parse

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var DirCmd = &cobra.Command{
	Use:   "dir",
	Short: "Parse a folder to the repo",
	Long:  "Parse a folder and record all the file(s) status to the repo",
	Run: func(cmd *cobra.Command, args []string) {
		log.Error("need implementation")
	},
}
