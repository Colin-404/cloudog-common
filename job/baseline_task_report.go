package job

type BaselineReport struct {
	TaskID    string          `json:"task_id"`
	AgentID   string          `json:"agent_id"`
	JobID     string          `json:"job_id"`
	Timestamp string          `json:"timestamp"`
	Summary   BaselineSummary `json:"summary"`
	Checks    []BaselineCheck `json:"checks"`
}

type BaselineSummary struct {
	TotalChecks  int `json:"total_checks"`
	PassedChecks int `json:"passed_checks"`
	FailedChecks int `json:"failed_checks"`
}

type BaselineCheck struct {
	ID          string                 `json:"id"`
	AgentID     string                 `json:"agent_id"`
	Title       string                 `json:"title"`
	Category    string                 `json:"category"`
	Severity    string                 `json:"severity"`
	Status      string                 `json:"status"`
	Scored      bool                   `json:"scored"`
	Description string                 `json:"description"`
	Remediation string                 `json:"remediation"`
	Impact      string                 `json:"impact"`
	Details     string                 `json:"details,omitempty"`
	References  []string               `json:"references,omitempty"`
	ResultData  map[string]interface{} `json:"result_data,omitempty"`
}
