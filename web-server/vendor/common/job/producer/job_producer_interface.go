package producer

import (
	"common/models/endpoint"
	"common/models/job"
)

/*
任务采用生产者消费者模式
这个是任务的生产者逻辑
*/

type JobProducerInterface interface {
	// GenerateCmd 组装，生成命令
	GenerateCmd(processorId string, end *endpoint.EndpointModel, target string, params interface{}) *job.JobCmd

	// PushCmd 存储命令
	PushCmd(cmd *job.JobCmd) *job.JobCommit

	// PopCmd 获取一个待执行的命令
	PopCmd(endpoint *endpoint.EndpointModel) *job.JobCmd

	// GetKey 获取命令的key
	GetKey(endpoint *endpoint.EndpointModel) string
}
