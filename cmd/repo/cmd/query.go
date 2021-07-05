package cmd

import (
	"github.com/spf13/cobra"
)

var queryCmd = &cobra.Command{
	Use:   "query",
	Short: "Query the repo",
	Long:  "Query all the records from repo by specified fields",
}
