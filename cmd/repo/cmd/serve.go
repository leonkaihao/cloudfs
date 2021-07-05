package cmd

import (
	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Create a web service",
	Long:  "Create a web UI service and listen",
}
