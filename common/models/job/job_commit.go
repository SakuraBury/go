package job

import (
	"common/models"
	"common/models/endpoint"
)

/*
任务返回的数据对象
*/

type JobCommit struct {
	/*
		命令ID，全局唯一，uuid
	*/
	CmdId string
	/*
		命令执行者信息
	*/
	Endpoint *endpoint.EndpointModel

	/*
		命令执行的返回结果
	*/
	Result models.BaseResult
}
