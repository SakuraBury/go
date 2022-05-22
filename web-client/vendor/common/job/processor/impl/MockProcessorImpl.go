package impl

import (
	"common/job/processor"
	"common/models/job"
	"fmt"
)

type MockProcessorImpl struct {
	processor.JobProcessorInterface
}

func (p *MockProcessorImpl) GetId() string {
	return "MOCK_PROCESSOR"
}
func (p *MockProcessorImpl) Process(cmd *job.JobCmd) *job.JobCommit {
	fmt.Println("=====mock process cmd:", cmd)
	rsp := job.JobCommit{}
	rsp.CmdId = cmd.CmdId
	rsp.Endpoint = cmd.Endpoint
	rsp.Result.Msg = "success"

	return &rsp
}

func init() {
	var inf processor.JobProcessorInterface
	p := MockProcessorImpl{}
	inf = &p
	factor := processor.GetJobProcessFactory()
	factor.RegisterProcessor(inf)
}
