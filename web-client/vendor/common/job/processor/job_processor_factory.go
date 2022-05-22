package processor

import "fmt"

type JobProcessFactory struct {
	processorMap map[string]JobProcessorInterface
}

// 将处理器注册进来
func (f *JobProcessFactory) RegisterProcessor(processor JobProcessorInterface) {
	key := processor.GetId()
	if "" == key {
		fmt.Errorf("job process not impl GetId\n")
		return
	}
	f.processorMap[processor.GetId()] = processor
}

// 根据处理器ID，获取处理器
func (f *JobProcessFactory) GetProcessor(processorId string) JobProcessorInterface {
	return f.processorMap[processorId]
}

var factory *JobProcessFactory

func init() {
	factory = new(JobProcessFactory)
	factory.processorMap = make(map[string]JobProcessorInterface)
}

func GetJobProcessFactory() *JobProcessFactory {
	return factory
}
