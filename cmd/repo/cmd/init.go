package cmd

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Init the repo",
	Long:  "Init the repo within a hidden .cfs folder in current directory",
	Run: func(cmd *cobra.Command, args []string) {
		log.Error("need implementation")
	},
}
