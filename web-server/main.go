package main

import (
	"github.com/beego/beego/v2/core/logs"
	"open-center/filter"
	_ "open-center/routers"

	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	logs.SetLogger(logs.AdapterMultiFile, `{"filename":"./logs/test.log","separate":["emergency", "alert", "critical", "error", "warning", "notice", "info", "debug"]}`)

	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.InsertFilter("*", beego.BeforeRouter, filter.AuthFilter)
	beego.Run()
}
