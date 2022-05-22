package endpoint

/**
endpoint 节点对象
*/
type EndpointModel struct {
	/*
		终端ID
	*/
	Id int64
	/*
		终端类型：NODE - open node;SERVICE - open service
	*/
	EndpointType string
	/*
		终端IP
	*/
	Ip string
	/*
		最近上报时间
	*/
	LastReportTime int64
}
