package job

import (
	"time"

	"github.com/colin-404/cloudog-common/proto"
)

// BaseJob 提供了 Job 接口的基本实现

type Job struct {
	JobID      string            `json:"job_id"`
	CreateTime string            `json:"create_time"`
	JobStatus  string            `json:"job_status"`
	JobType    string            `json:"job_type"`
	JobParams  map[string]string `json:"job_params"`
	Tasks      []*proto.Task     `json:"task"`
	Summary    interface{}       `json:"summary"`
	Results    interface{}       `json:"results"`
}

// 创建新的 BaselineJob 的构造函数
func NewJob(jobType string, params map[string]string) *Job {
	return &Job{
		JobID:      GenerateJobID(),
		JobType:    jobType,
		JobParams:  params,
		CreateTime: time.Now().Format(time.RFC3339),
	}
}
