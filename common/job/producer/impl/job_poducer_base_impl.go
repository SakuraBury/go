package impl

import (
	"common/job/producer"
	"common/models/endpoint"
	"common/models/job"
	"strconv"
)

type BaseJobProducer struct {
	producer.JobProducerInterface
}

// PushCmd
// 默认情况下把命令丢到一个kv存储中
// key 是终端类型+ip
// value 是cmd对象list
func (p *BaseJobProducer) PushCmd(cmd *job.JobCmd) *job.JobCommit {

	return nil
}

func (p *BaseJobProducer) GetKey(endpoint *endpoint.EndpointModel) string {
	return strconv.FormatInt(endpoint.Id, 10)
}
