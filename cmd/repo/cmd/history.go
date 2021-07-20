package cmd

import (
	"context"
	"fmt"
	"math"
	"os"
	"strconv"
	"time"

	"github.com/leonkaihao/cloudfs/pkg/repository"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var historyCmd = &cobra.Command{
	Use:   "history",
	Short: "User actions history",
	Long:  "User actions history",
	Run: func(cmd *cobra.Command, args []string) {
		var (
			path string
			err  error
		)
		if path, err = os.Getwd(); err != nil {
			log.Fatal(err)
		}
		repo := repository.New(path)
		if err = repo.Load(); err != nil {
			log.Error(err)
			return
		}

		var (
			maxCount  int
			actionsTb repository.Actions
			ok        bool
		)
		maxCount = math.MaxInt32
		if len(args) == 1 {
			maxCount, err = strconv.Atoi(args[0])
			if err != nil {
				log.Fatal(err)
			}
		}
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()
		if actionsTb, ok = repo.Table(repository.ActionsTableName).(repository.Actions); !ok {
			log.Error("History data is invalid.")
			return
		}
		actions, err := actionsTb.Items(ctx, maxCount)
		if err != nil {
			log.Error(err)
			return
		}
		for _, action := range actions {
			fmt.Println(action)
		}
	},
}
