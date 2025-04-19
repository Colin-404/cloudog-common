package comm

import "github.com/colin-404/cloudog-common/mproto"

type Job struct {
	ID       string            `json:"id"`
	Status   string            `json:"status"`
	Params   map[string]string `json:"params"`
	TaskList []mproto.Task     `json:"task_list"`
}

type JobRequest struct {
	Type   string            `json:"type"`
	Params map[string]string `json:"params"`
}
