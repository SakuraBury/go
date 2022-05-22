package schedule

import (
	"bytes"
	"common/job/processor"
	"common/request"
	"common/response"
	"encoding/json"
	"fmt"
	"github.com/robfig/cron/v3"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

const CODE_CENTER_HEART_ADDR = "http://127.0.0.1:8080/v1/heartbeat/report"

func init() {
	crontab := cron.New(cron.WithSeconds())

	task := func() {
		fmt.Println("ssss...")
		rsp := heartbeat()
		fmt.Printf("heartbeat rsp data: %v", rsp)
		if nil != rsp.Cmd { // 这就表示有命令要执行
			prcoService := processor.JobProcessService{}
			prcoService.Process(rsp.Cmd)
		}
	}

	spec := "*/10 * * * * ?"
	crontab.AddFunc(spec, task)
	crontab.Start()
}

func heartbeat() *response.HeartbeatResponse {
	ip := getinternal()
	if "" == ip {
		fmt.Println("not get ip.....")
		return nil
	}
	req := request.HeartbeatRequest{}
	req.Ip = ip
	req.EndpointType = "SERVICE"
	req.ReportTime = time.Time{}.UnixMilli()

	bReq, err := json.Marshal(req)
	if err != nil {
		fmt.Println("json err")
		return nil
	}

	rsp, err := http.Post(CODE_CENTER_HEART_ADDR, "application/json", bytes.NewReader(bReq))
	if err != nil {
		fmt.Println("heart err")
		return nil
	}

	body, _ := ioutil.ReadAll(rsp.Body)

	rspData := response.HeartbeatResponse{}

	json.Unmarshal(body, &rspData)

	return &rspData
}

func getinternal() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}
