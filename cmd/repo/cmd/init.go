package cmd

import (
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/leonkaihao/cloudfs/pkg/repository"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Init the repo",
	Long:  "Init the repo within a hidden .cfs folder in current directory",
	Run: func(cmd *cobra.Command, args []string) {
		var (
			path string
			err  error
		)
		if path, err = os.Getwd(); err != nil {
			log.Fatal(err)
		}
		repo := repository.New(path)
		if err = repo.Init(); err != nil {
			log.Error(err)
		}
	},
}
