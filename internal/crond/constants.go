package crond

import (
	"github.com/hibiken/asynq"
)

type TaskType = string

const (
	TaskTypeDebug TaskType = "DebugTask"
)

type QueueType = string

const (
	QueueTypeDefault QueueType = "default"
)

func getQueue(t TaskType) QueueType {
	switch t {
	default:
		return QueueTypeDefault
	}
}

func getOps(t TaskType) []asynq.Option {
	switch t {
	default:
		return []asynq.Option{}
	}
}
