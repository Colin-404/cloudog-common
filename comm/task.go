package comm

const (
	SyncConfigTask    = 1000
	PluginTask        = 3000
	LinuxBaselineTask = 3001
)

type PluginConfig struct {
	Name         string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Type         string   `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	Version      string   `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
	Sha256       string   `protobuf:"bytes,4,opt,name=sha256,proto3" json:"sha256,omitempty"`
	Signature    string   `protobuf:"bytes,5,opt,name=signature,proto3" json:"signature,omitempty"`
	DownloadUrls []string `protobuf:"bytes,6,rep,name=download_urls,json=downloadUrls,proto3" json:"download_urls,omitempty"`
	Detail       string   `protobuf:"bytes,7,opt,name=detail,proto3" json:"detail,omitempty"`
}

type AgentConfig struct {
	AgentID       string          `json:"agent_id"`
	AgentVersion  string          `json:"agent_version"`
	PluginConfigs []*PluginConfig `json:"plugin_configs"`
}

type TaskData struct {
	TaskID     string      `json:"task_id"`
	TaskName   string      `json:"task_name"`
	TaskType   string      `json:"task_type"`
	TaskStatus string      `json:"task_status"`
	TaskResult interface{} `json:"task_result"`
}
