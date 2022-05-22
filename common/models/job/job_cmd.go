package job

import (
	"common/models/endpoint"
	"time"
)

/*
任务命令参数对象
*/

type JobCmd struct {
	/*
		命令ID，全局唯一，uuid
	*/
	CmdId string
	/*
		命令执行者信息
	*/
	Endpoint *endpoint.EndpointModel
	/*
		命令字，如shell的 echo 命令
	*/

	/*
		处理器ID
	*/
	ProcessorId string

	/*
		命令目标，如:echo， ./xxx
	*/
	Target string
	/*
		命令的参数
	*/
	Params interface{}

	/*
		是否是异步执行
	*/
	IsAsync bool

	/*
		是否可重入
	*/
	CanRetry bool

	/*
		命令创建时间
	*/
	CreateTime time.Time
	/*
		最后一次发送时间
	*/
	LastSendTime time.Time

	/*
		发送次数
	*/
	SendNum int
}
