package beat

import (
	"context"
	"log"

	"github.com/hibiken/asynq"
	"github.com/ovinc-cn/apicenter/v2/internal/crond"
	"github.com/ovinc-cn/apicenter/v2/pkg/configUtils"
)

type taskConfig struct {
	TaskType crond.TaskType
	Data     interface{}
	CronSpec string
	Ops      []asynq.Option
}

var periodTasks = []taskConfig{
	{
		TaskType: crond.TaskTypeDebug,
		CronSpec: "* * * * *",
	},
}

func Serve() {
	// init beat
	scheduler := asynq.NewScheduler(crond.RedisClientOption, &asynq.SchedulerOpts{Location: configUtils.AppTimezone()})

	// register task
	for _, periodTask := range periodTasks {
		// register task
		task, err := crond.NewTask(context.Background(), periodTask.TaskType, periodTask.Data, periodTask.Ops...)
		if err != nil {
			log.Fatalf("[Beat] create %s task failed: %s", periodTask.TaskType, err)
		}
		if _, err = scheduler.Register(periodTask.CronSpec, task); err != nil {
			log.Fatalf("[Beat] register %s task failed: %s", periodTask.TaskType, err)
		}
		log.Printf("[Beat] register %s task success", periodTask.TaskType)
	}

	// start
	if err := scheduler.Run(); err != nil {
		log.Fatalf("[Beat] start beat failed: %s", err)
	}
}
