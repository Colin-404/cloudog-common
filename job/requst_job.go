package job

type JobCreateRequest struct {
	JobType   string            `json:"job_type"`
	JobParams map[string]string `json:"job_params"`
}
