package main

import (
	_ "common/job/processor"
	_ "common/job/processor/impl"
	"github.com/beego/beego/v2/core/logs"

	beego "github.com/beego/beego/v2/server/web"
	_ "open-service/schedule"
)

func main() {
	logs.SetLogger(logs.AdapterMultiFile, `{"filename":"./logs/test.log","separate":["emergency", "alert", "critical", "error", "warning", "notice", "info", "debug"]}`)
	beego.Run()
}
