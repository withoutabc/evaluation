package model

type YarnApplication struct {
	Id                     string `json:"id"` //job id
	User                   string
	Name                   string `json:"name"`        //job名称
	State                  string `json:"state"`       //是否完成
	FinalStatus            string `json:"finalStatus"` //成功还是失败
	Progress               int32  `json:"progress"`    //进度0-100
	TrackingUI             string
	TrackingUrl            string `json:"trackingUrl"` //追踪url
	Diagnostics            string
	ClusterId              int
	ApplicationType        string
	ApplicationTags        string
	Priority               int
	StartedTime            int64 `json:"startedTime"`  //开始时间
	LaunchTime             int64 `json:"launchTime"`   //启动时间
	FinishedTime           int64 `json:"finishedTime"` //完成时间
	ElapsedTime            int   `json:"elapsedTime"`  //用时
	AmContainerLogs        string
	AmHostHttpAddress      string
	AmRPCAddress           string
	MasterNodeId           string
	AllocatedMB            int
	AllocatedVCores        int
	ReservedMB             int
	ReservedVCores         int
	RunningContainers      int
	MemorySeconds          int
	VcoreSeconds           int
	QueueUsagePercentage   int
	ClusterUsagePercentage int
	ResourceSecondsMap     struct {
		Entry struct {
			Key   string
			Value string
		}
	}
	PreemptedResourceMB         int
	PreemptedResourceVCores     int
	NumNonAMContainerPreempted  int
	NumAMContainerPreempted     int
	PreemptedMemorySeconds      int
	PreemptedVcoreSeconds       int
	PreemptedResourceSecondsMap struct{}
	LogAggregationStatus        string
	UnmanagedApplication        bool
	AmNodeLabelExpression       string
	Timeouts                    struct {
		Timeout struct {
			Type                   string
			ExpiryTime             string
			RemainingTimeInSeconds int
		}
	}
}
type YarnApplications struct {
	Apps struct {
		App []YarnApplication `json:"app"`
	} `json:"apps"`
}
