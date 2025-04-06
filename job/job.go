package job

import (
	"time"

	"github.com/colin-404/cloudog-common/proto"
)

// BaseJob 提供了 Job 接口的基本实现
type Job struct {
	JobID      string            `json:"job_id"`
	CreateTime string            `json:"create_time"`
	JobType    string            `json:"job_type"`
	JobParams  map[string]string `json:"job_params"`
	Tasks      []*proto.Task     `json:"task"`
	Summary    interface{}       `json:"summary"`
	Results    interface{}       `json:"results"`
}

// 实现 Job 接口的方法
func (j *Job) GetJobID() string {
	return j.JobID
}

func (j *Job) GetJobType() string {
	return j.JobType
}

func (j *Job) GetJobParams() map[string]string {
	return j.JobParams
}

func (j *Job) GetSummary() interface{} {
	return j.Summary
}

func (j *Job) GetResults() interface{} {
	return j.Results
}

func (j *Job) SetSummary(summary interface{}) {
	j.Summary = summary
}

func (j *Job) SetResults(results interface{}) {
	j.Results = results
}

// 创建新的 BaselineJob 的构造函数
func (j *Job) NewJob(jobType string, params map[string]string) *Job {
	return &Job{
		JobID:      GenerateJobID(),
		JobType:    jobType,
		JobParams:  params,
		CreateTime: time.Now().Format(time.RFC3339),
	}
}
