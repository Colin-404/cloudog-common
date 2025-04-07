package job

import (
	"github.com/colin-404/cloudog-common/proto"
)

func (j *Job) NewBaselineJob(params map[string]string) (*Job, error) {
	jobType := "baseline"
	job := j.NewJob(jobType, params)
	job.Summary = &JobBaselineSummary{}
	job.Results = &JobBaselineChecks{}
	return job, nil
}

// type BaselineJob struct {
// 	JobID     string              `json:"job_id"`
// 	JobType   string              `json:"job_type"`
// 	JobParams map[string]string   `json:"job_params"`
// 	Timestamp string              `json:"timestamp"`
// 	Tasks     []proto.Task        `json:"tasks"`
// 	Summary   JobBaselineSummary  `json:"summary"`
// 	Checks    []JobBaselineChecks `json:"checks"`
// }

type JobBaselineSummary struct {
	TotalChecks      int               `json:"total_checks"`
	PassedChecks     int               `json:"passed_checks"`
	FailedChecks     int               `json:"failed_checks"`
	BaselineSummarys []BaselineSummary `json:"baseline_summary"`
}

type JobBaselineChecks struct {
	Agent  proto.Agent     `json:"agent"`
	Checks []BaselineCheck `json:"checks"`
}
