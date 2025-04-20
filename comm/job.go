package comm

import "github.com/colin-404/cloudog-common/mproto"

const (
	JobStatusPending  = "pending"
	JobStatusRunning  = "running"
	JobStatusSuccess  = "success"
	JobStatusFailed   = "failed"
	LinuxBaselineJob  = 500
	AWSBaselineJob    = 501
	AliyunBaselineJob = 502
	KubeBaselineJob   = 503
)

type Job struct {
	ID       string            `json:"id"`
	Status   string            `json:"status"`
	Type     int32             `json:"type"`
	Params   map[string]string `json:"params"`
	TaskList []mproto.Task     `json:"task_list"`
}
