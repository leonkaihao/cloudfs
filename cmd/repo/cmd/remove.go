package cmd

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "remove file records",
	Long:  "remove file records from repo",
	Run: func(cmd *cobra.Command, args []string) {
		log.Error("need implementation")
	},
}
