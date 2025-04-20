package comm

const (
	BaselineData            = 8000
	BaselineDataStatus      = 8001
	LinuxBaselinePluginName = "linux_baseline"
	//debian/ubuntu
	DebianLinuxBaselineID = 1300
	//centos/redhat
	CentosLinuxBaselineID = 1301
)

type BaselineInfo struct {
	BaselineId      int         `json:"baseline_id" bson:"baseline_id"`
	BaselineVersion string      `json:"baseline_version" bson:"baseline_version"`
	Status          string      `json:"status" bson:"status"`
	Msg             string      `json:"msg" bson:"msg"`
	CheckList       []CheckInfo `json:"check_list" bson:"check_list"`
}

type CheckInfo struct {
	CheckId       int    `json:"check_id" bson:"check_id"`
	Severity      string `json:"security" bson:"security"`
	Type          string `json:"type" bson:"type"`
	Title         string `json:"title" bson:"title"`
	Description   string `json:"description" bson:"description"`
	Solution      string `json:"solution" bson:"solution"`
	TypeCn        string `json:"type_cn" bson:"type_cn"`
	TitleCn       string `json:"title_cn" bson:"title_cn"`
	DescriptionCn string `json:"description_cn" bson:"description_cn"`
	SolutionCn    string `json:"solution_cn" bson:"solution_cn"`
	Result        int    `json:"result" bson:"result"`
	Msg           string `json:"msg" bson:"msg"`
}
