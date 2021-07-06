package cmd

import (
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/leonkaihao/cloudfs/pkg/services"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Init the repo",
	Long:  "Init the repo within a hidden .cfs folder in current directory",
	Run: func(cmd *cobra.Command, args []string) {
		svc := services.New()
		var (
			path string
			err  error
		)
		if path, err = os.Getwd(); err != nil {
			log.Fatal(err)
		}
		if err = svc.Init(path); err != nil {
			log.Fatal(err)
		}
	},
}
