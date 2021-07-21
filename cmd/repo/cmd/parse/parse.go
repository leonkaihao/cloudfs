package parse

import (
	"github.com/spf13/cobra"
)

var ParseCmd = &cobra.Command{
	Use:   "parse",
	Short: "Parse a file/folder to the repo",
	Long:  "Parse a file/folder and record all the file(s) status to the repo",
}

func init() {
	ParseCmd.AddCommand(DirCmd)
	ParseCmd.AddCommand(FileCmd)
}
