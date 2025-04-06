package job

type JobCreateRequest struct {
	JobType   string            `json:"job_type"`
	JobParams map[string]string `json:"job_params"`
}

type JobResultRequest struct {
	Time      string            `json:"time"`
	JobType   string            `json:"job_type"`
	JobParams map[string]string `json:"job_params"`
}
