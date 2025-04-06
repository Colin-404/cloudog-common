package job

// BaseJob 提供了 Job 接口的基本实现
type Job struct {
	JobID      string            `json:"job_id"`
	CreateTime string            `json:"create_time"`
	JobType    string            `json:"job_type"`
	JobParams  map[string]string `json:"job_params"`
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
