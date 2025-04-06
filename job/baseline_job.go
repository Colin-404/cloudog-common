package job

import (
	"fmt"

	"github.com/colin-404/cloudog-common/proto"
)

func NewBaselineJob(params JobCreateRequest) (*Job, error) {
	if params.JobType != "baseline" {
		return nil, fmt.Errorf("invalid job type: %s", params.JobType)
	}
	job := NewJob(params.JobType, params.JobParams)
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
