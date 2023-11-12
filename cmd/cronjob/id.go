package cronjob

import (
	"context"

	"github.com/go-co-op/gocron"
	"github.com/sociosarbis/grpc/boilerplate/internal/config"
	"github.com/sociosarbis/grpc/boilerplate/internal/goredis"
	"github.com/sociosarbis/grpc/boilerplate/internal/pkg/errgo"
)

func NewIDCronJob(cfg config.AppConfig, clients *goredis.RedisClients, scheduler *gocron.Scheduler, jobs registeredJobs) error {

	if _, ok := jobs["expireWorkerID"]; ok {
		_, err := scheduler.Every(30).Second().Do(func() {
			clients.ExpireWorkerID(context.Background(), clients.Main, "service_worker")
		})

		if err != nil {
			return errgo.Wrap(err, "clients.ExpireWorkerID")
		}
	}
	return nil
}
