package job

import (
	"time"

	"github.com/colin-404/cloudog-common/proto"
)

// BaseJob 提供了 Job 接口的基本实现

type JobReport struct {
	JobID      string            `json:"job_id"`
	CreateTime string            `json:"create_time"`
	JobStatus  string            `json:"job_status"`
	JobType    string            `json:"job_type"`
	JobParams  map[string]string `json:"job_params"`
	Tasks      []*proto.Task     `json:"task"`
	Summary    interface{}       `json:"summary"`
	Results    interface{}       `json:"results"`
}

// 实现 Job 接口的方法
func (j *JobReport) GetJobID() string {
	return j.JobID
}

func (j *JobReport) GetJobType() string {
	return j.JobType
}

func (j *JobReport) GetJobParams() map[string]string {
	return j.JobParams
}

func (j *JobReport) GetSummary() interface{} {
	return j.Summary
}

func (j *JobReport) GetResults() interface{} {
	return j.Results
}

func (j *JobReport) SetSummary(summary interface{}) {
	j.Summary = summary
}

func (j *JobReport) SetResults(results interface{}) {
	j.Results = results
}

// 创建新的 BaselineJob 的构造函数
func (j *JobReport) NewJobReport(jobType string, params map[string]string) *JobReport {

	return &JobReport{
		JobID:      GenerateJobID(),
		JobType:    jobType,
		JobParams:  params,
		CreateTime: time.Now().Format(time.RFC3339),
	}
}
