package query

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var HashCmd = &cobra.Command{
	Use:   "hash",
	Short: "query files with specific hash",
	Long:  "query files with specific hash from repo",
	Run: func(cmd *cobra.Command, args []string) {
		log.Error("need implementation")
	},
}
