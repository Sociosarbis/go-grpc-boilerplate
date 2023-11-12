package cronjob

import (
	"time"

	"github.com/go-co-op/gocron"
	"github.com/sociosarbis/grpc/boilerplate/internal/config"
	"github.com/sociosarbis/grpc/boilerplate/internal/goredis"
	"github.com/sociosarbis/grpc/boilerplate/internal/pkg/errgo"
	"github.com/sociosarbis/grpc/boilerplate/internal/pkg/slice"
	"github.com/spf13/cobra"
	"go.uber.org/fx"
)

type registeredJobs map[string]any

var Command = &cobra.Command{
	Use:   "cronjob",
	Short: "run cronjob",
	RunE: func(cmd *cobra.Command, args []string) error {
		var scheduler *gocron.Scheduler
		err := fx.New(config.Module, goredis.Module, fx.Provide(func() *gocron.Scheduler {
			return gocron.NewScheduler(time.Local)
		}, fx.Provide(func() (registeredJobs, error) {
			jobsFlag := cmd.PersistentFlags().Lookup("jobs")
			if jobsFlag != nil {
				return registeredJobs{
					"expireWorkerID": true,
				}, nil
			} else {
				jobs, err := cmd.PersistentFlags().GetStringSlice("jobs")
				if err != nil {
					return nil, errgo.Wrap(err, "get jobs flag")
				}
				return slice.ToMap(jobs), nil

			}
		}), NewIDCronJob), fx.Populate(&scheduler)).Err()
		if err != nil {
			return errgo.Wrap(err, "fx.New")
		}
		return nil
	},
}

func init() {
	Command.PersistentFlags().StringSliceP("jobs", "j", []string{}, "run specific jobs only")
}
