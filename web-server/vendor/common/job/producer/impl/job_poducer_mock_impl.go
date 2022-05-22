package impl

import (
	"common/job/producer"
	"common/models/endpoint"
	"common/models/job"
	"github.com/beego/beego/v2/core/logs"
	"github.com/google/uuid"
	"strconv"
	"time"
)

type MockJobProducer struct {
	producer.JobProducerInterface

	cmdMap map[string][]*job.JobCmd
}

func (p *MockJobProducer) GenerateCmd(processorId string, end *endpoint.EndpointModel, target string, params interface{}) *job.JobCmd {
	cmd := job.JobCmd{}
	cmd.Endpoint = end
	cmd.Target = target
	cmd.Params = params
	cmd.CreateTime = time.Time{}
	cmd.CmdId = uuid.NewString()
	cmd.ProcessorId = processorId
	return &cmd
}

// PushCmd
// mock 时直接存本地map
// key 是终端 id
// value 是cmd对象list
func (p *MockJobProducer) PushCmd(cmd *job.JobCmd) *job.JobCommit {

	if p.cmdMap == nil {
		p.cmdMap = make(map[string][]*job.JobCmd)
	}
	key := p.GetKey(cmd.Endpoint)
	cmds := p.cmdMap[key]

	rsp := job.JobCommit{}
	rsp.CmdId = cmd.CmdId
	rsp.Endpoint = cmd.Endpoint

	for _, value := range cmds {
		if value.CmdId == cmd.CmdId {
			logs.Debug("cmd has exist")
			rsp.Result.Code = -1
			rsp.Result.Msg = "cmd has exist"
			return &rsp
		}
	}
	cmds = append(cmds, cmd)

	p.cmdMap[key] = cmds
	return &rsp
}

func (p *MockJobProducer) PopCmd(endpoint *endpoint.EndpointModel) *job.JobCmd {
	if nil == p.cmdMap {
		return nil
	}
	key := p.GetKey(endpoint)

	cmds := p.cmdMap[key]
	if nil == cmds || len(cmds) == 0 {
		return nil
	}

	rsp := cmds[0]
	rsp.SendNum++
	rsp.LastSendTime = time.Time{}
	return rsp
}

func (p *MockJobProducer) GetKey(endpoint *endpoint.EndpointModel) string {
	return strconv.FormatInt(endpoint.Id, 10)
}
