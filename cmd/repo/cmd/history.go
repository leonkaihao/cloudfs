package cmd

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var historyCmd = &cobra.Command{
	Use:   "history",
	Short: "User actions history",
	Long:  "User actions history",
	Run: func(cmd *cobra.Command, args []string) {
		log.Error("need implementation")
	},
}
