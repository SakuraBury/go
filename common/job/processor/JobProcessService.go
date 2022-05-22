package processor

import (
	"common/models"
	"common/models/job"
	"fmt"
	"github.com/beego/beego/v2/core/logs"
)

type JobProcessService struct {
	cmdLock map[string]int

	commitCache map[string]*job.JobCommit
}

func (service *JobProcessService) Process(cmd *job.JobCmd) *job.JobCommit {
	rsp := job.JobCommit{
		CmdId:    cmd.CmdId,
		Endpoint: cmd.Endpoint,
		Result:   models.BaseResult{},
	}
	// 第一步：根据命令的处理器ID，找到处理器
	processor := GetJobProcessFactory().GetProcessor(cmd.ProcessorId)
	if nil == processor {
		logs.Error("prcessor id:%s not exist\n", cmd.ProcessorId)
		rsp.Result.Code = -100
		rsp.Result.Msg = "processor id not exist"
		return &rsp
	}
	// 第二步：加锁
	// TODO: 后面用redis
	if service.cmdLock[cmd.CmdId] == 1 {
		logs.Debug("===is processing cmd:", cmd.CmdId)
		rsp.Result.Code = 1000
		rsp.Result.Msg = "cmd is processing"
		return &rsp
	}

	// 第三步：如果是异步任务，先看缓存是否有结果，有就返回
	// TODO: 这个缓存后面也要有redis
	if cmd.IsAsync {
		cache := service.GetCommitCache(cmd.CmdId)
		if nil != cache {
			fmt.Println("===is processing cache cmd:", cmd.CmdId)
			return cache
		}
	}

	// 第四步：执行任务,如果是异步任务，需要将结果缓存本地
	if cmd.IsAsync {
		go processor.Process(cmd)
		rsp.Result.Code = 10000
		rsp.Result.Msg = "INIT_PROCESSING"
		service.CacheCommit(cmd.CmdId, &rsp, 3600)
	} else {
		rsp = *processor.Process(cmd)
	}
	return &rsp
}

// TODO: 这个缓存后面也要有redis
func (service *JobProcessService) GetCommitCache(cmdId string) *job.JobCommit {
	return service.commitCache[cmdId]
}

// TODO: 这个后面也有用redis
func (service *JobProcessService) CacheCommit(cmdId string, commit *job.JobCommit, expire int64) {
	service.commitCache[cmdId] = commit
}
