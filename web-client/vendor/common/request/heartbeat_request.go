package request

type HeartbeatRequest struct {
	/*
		终端类型：NODE - open node;SERVICE - open service
	*/
	EndpointType string
	/*
		上报IP
	*/
	Ip string
	/*
		上报时间
	*/
	ReportTime int64
}
