package history

import (
	"context"
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	cleanFromID int
	cleanToID   int
)

var CleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "clean history",
	Long:  "clean a specified item by ID or clean all records",
	Run: func(cmd *cobra.Command, args []string) {
		repo, actionsTb, err := initHistory()
		if err != nil {
			log.Error(err)
			return
		}
		defer repo.Close()
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()
		if err = actionsTb.Clean(ctx, cleanFromID, cleanToID); err != nil {
			log.Error(err)
			return
		}
		fmt.Println("history is cleaned.")
	},
}
