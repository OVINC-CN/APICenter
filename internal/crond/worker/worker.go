package worker

import (
	"log"

	"github.com/hibiken/asynq"
	"github.com/ovinc-cn/apicenter/v2/internal/apps/home"
	"github.com/ovinc-cn/apicenter/v2/internal/config"
	"github.com/ovinc-cn/apicenter/v2/internal/crond"
	"github.com/ovinc-cn/apicenter/v2/pkg/asyncTask"
)

func Serve() {
	// init worker
	srv := asynq.NewServer(
		crond.RedisClientOption,
		asynq.Config{
			Concurrency:     config.Config.Cron.WorkerConcurrency,
			Queues:          config.Config.Cron.WorkerQueues,
			ShutdownTimeout: config.Config.Cron.WorkerShutDownTimeout,
		},
	)

	// init
	mux := asynq.NewServeMux()
	mux.Use(asyncTask.RecoverMiddleware, asyncTask.ObserveMiddleware)

	// home
	mux.HandleFunc(crond.TaskTypeDebug, home.DebugTask)

	// run worker
	if err := srv.Run(mux); err != nil {
		log.Fatalf("[Worker] start failed: %s", err)
	}
}
