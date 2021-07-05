package query

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var PathCmd = &cobra.Command{
	Use:   "path",
	Short: "query path",
	Long:  "query path from repo",
	Run: func(cmd *cobra.Command, args []string) {
		log.Error("need implementation")
	},
}
