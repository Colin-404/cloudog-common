package comm

const (
	HeartBeatData       = 2000
	PluginHeartBeatData = 2001
)

type Heartbeat struct {
	KernelVersion   string `json:"kernel_version"`
	Arch            string `json:"arch"`
	Platform        string `json:"platform"`
	PlatformFamily  string `json:"platform_family"`
	PlatformVersion string `json:"platform_version"`
	State           string `json:"state"`
	StateDetail     string `json:"state_detail"`
	NetMode         string `json:"net_mode"`
	RxSpeed         string `json:"rx_speed"`
	TxSpeed         string `json:"tx_speed"`
	CPU             string `json:"cpu"`
	RSS             string `json:"rss"`
	ReadSpeed       string `json:"read_speed"`
	WriteSpeed      string `json:"write_speed"`
	PID             string `json:"pid"`
	NFD             string `json:"nfd"`
	StartTime       string `json:"start_time"`
	TxTPS           string `json:"tx_tps"`
	RxTPS           string `json:"rx_tps"`
	Du              string `json:"du"`
	NGR             string `json:"ngr"`
	NProc           string `json:"nproc"`
	TotalMem        string `json:"total_mem"`
	Load1           string `json:"load_1"`
	Load5           string `json:"load_5"`
	Load15          string `json:"load_15"`
	RunningProcs    string `json:"running_procs"`
	TotalProcs      string `json:"total_procs"`
	HostSerial      string `json:"host_serial"`
	HostID          string `json:"host_id"`
	HostModel       string `json:"host_model"`
	HostVendor      string `json:"host_vendor"`
	DNS             string `json:"dns"`
	Gateway         string `json:"gateway"`
	CPUName         string `json:"cpu_name"`
	BootTime        string `json:"boot_time"`
	CPUUsage        string `json:"cpu_usage"`
	MemUsage        string `json:"mem_usage"`
}

type PluginHeartbeat struct {
	Name       string `json:"name"`
	Version    string `json:"version"`
	CPU        string `json:"cpu"`
	RSS        string `json:"rss"`
	ReadSpeed  string `json:"read_speed"`
	WriteSpeed string `json:"write_speed"`
	PID        string `json:"pid"`
	NFD        string `json:"nfd"`
	StartTime  string `json:"start_time"`
	Du         string `json:"du"`
	RxTPS      string `json:"rx_tps"`
	TxTPS      string `json:"tx_tps"`
	RxSpeed    string `json:"rx_speed"`
	TxSpeed    string `json:"tx_speed"`
}
