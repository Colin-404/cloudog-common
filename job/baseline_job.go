package job

import "github.com/colin-404/cloudog-common/proto"

func NewBaselineJob(params map[string]string) *Job {
	job := NewJob(params)
	job.JobType = "baseline"
	job.Summary = &JobBaselineSummary{}
	job.Results = &JobBaselineChecks{}
	return job
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
