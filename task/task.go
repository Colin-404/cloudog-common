package task

import (
	"time"

	"github.com/colin-404/cloudog-common/utils"
)

func NewTask(taskType string, taskParams map[string]string) *Task {
	return &Task{
		TaskId:     utils.GenerateTaskID(),
		CreateTime: time.Now().Format(time.RFC3339),
		TaskType:   taskType,
		TaskParams: taskParams,
	}
}
