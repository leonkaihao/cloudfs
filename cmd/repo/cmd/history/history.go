package history

import (
	"context"
	"errors"
	"fmt"
	"math"
	"os"
	"time"

	"github.com/leonkaihao/cloudfs/pkg/repository"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	historyMaxCount int
)

var HistoryCmd = &cobra.Command{
	Use:   "history",
	Short: "User actions history",
	Long:  "User actions history",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		_, actionsTb, err := initHistory()
		if err != nil {
			log.Error(err)
			return
		}
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()
		actions, err := actionsTb.Items(ctx, historyMaxCount)
		if err != nil {
			log.Error(err)
			return
		}
		for _, action := range actions {
			fmt.Println(action)
		}
	},
}

func init() {
	HistoryCmd.Flags().IntVarP(
		&historyMaxCount, "number", "n",
		math.MaxInt32, "The largest number of items that can be shown")
	HistoryCmd.AddCommand(CleanCmd)
	CleanCmd.Flags().IntVarP(
		&cleanFromID, "from", "f",
		0, "Clean history from ID")
	CleanCmd.Flags().IntVarP(
		&cleanToID, "to", "t",
		0, "Clean history to ID")
}

func initHistory() (repository.Service, repository.Actions, error) {
	var (
		path string
		err  error
	)
	if path, err = os.Getwd(); err != nil {
		return nil, nil, err
	}
	repo := repository.New(path)
	if err = repo.Load(); err != nil {
		return nil, nil, err
	}

	var (
		actionsTb repository.Actions
		ok        bool
	)
	if actionsTb, ok = repo.Table(repository.ActionsTableName).(repository.Actions); !ok {
		return nil, nil, errors.New("History data is invalid")
	}
	return repo, actionsTb, nil
}
