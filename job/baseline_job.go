package job

func (j *JobReport) NewBaselineJobReport(params map[string]string) (*JobReport, error) {
	jobType := "baseline"
	job := j.NewJobReport(jobType, params)
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
	AgentID string          `json:"agent_id"`
	Checks  []BaselineCheck `json:"checks"`
}
