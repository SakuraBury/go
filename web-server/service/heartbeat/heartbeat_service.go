package heartbeat

import (
	"common/models/endpoint"
	"common/models/job"
	"common/request"
	"fmt"
	"open-center/common"
)

/**
心跳管理服务
*/
var index int64 = 1

// 心跳上报
func Report(req *request.HeartbeatRequest) *job.JobCmd {
	fmt.Printf("heartbeat report req ip： %s", req.Ip)
	id := GetIdByIp(req.Ip, req.EndpointType)

	if id <= 0 {
		id = index
		index++
	}
	end := endpoint.EndpointModel{}
	end.Id = id
	end.Ip = req.Ip
	end.LastReportTime = req.ReportTime
	end.EndpointType = req.EndpointType
	SaveReport(&end)

	fmt.Println("heater end:", end)
	// 查看是否有任务
	producer := common.GetMockProducerInstance()

	return producer.PopCmd(&end)
}
