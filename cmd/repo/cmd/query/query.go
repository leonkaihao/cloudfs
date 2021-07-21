package query

import (
	"github.com/spf13/cobra"
)

var QueryCmd = &cobra.Command{
	Use:   "query",
	Short: "Query the repo",
	Long:  "Query all the records from repo by specified fields",
}

func init() {
	QueryCmd.AddCommand(HashCmd)
	QueryCmd.AddCommand(PathCmd)
}
