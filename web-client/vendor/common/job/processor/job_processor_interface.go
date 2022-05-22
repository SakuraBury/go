package processor

import "common/models/job"

/*
任务处理器接口
*/
type JobProcessorInterface interface {
	// GetId 定义处理器的ID
	GetId() string

	// Process 处理任务
	Process(cmd *job.JobCmd) *job.JobCommit
}
