package job

import (
	"time"

	"github.com/colin-404/cloudog-common/proto"
)

func NewTask(taskType string, taskParams map[string]string) *proto.Task {
	return &proto.Task{
		TaskId:     GenerateTaskID(),
		CreateTime: time.Now().Format(time.RFC3339),
		TaskType:   taskType,
		TaskParams: taskParams,
	}
}
