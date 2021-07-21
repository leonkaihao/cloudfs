package cmd

import (
	"github.com/leonkaihao/cloudfs/cmd/repo/cmd/history"
	"github.com/leonkaihao/cloudfs/cmd/repo/cmd/parse"
	"github.com/leonkaihao/cloudfs/cmd/repo/cmd/query"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "repo",
	Short: "repo is a tool for creating cloud fs repo",
	Long:  "A fast and flexible repo generator built in Go",
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(initCmd)
	rootCmd.AddCommand(history.HistoryCmd)
	rootCmd.AddCommand(parse.ParseCmd)
	rootCmd.AddCommand(query.QueryCmd)

	rootCmd.AddCommand(removeCmd)
	rootCmd.AddCommand(serveCmd)
	rootCmd.AddCommand(watchCmd)
}
