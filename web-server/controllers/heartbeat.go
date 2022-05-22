package controllers

import (
	"common/request"
	"common/response"
	"encoding/json"
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
	"open-center/service/heartbeat"
)

/**
open-service 和 open-endpoint 心跳上报服务
*/

type HeartbeatController struct {
	beego.Controller
}

func (u *HeartbeatController) Report() {
	logs.Debug("====report token:%s", u.Ctx.Input.GetData("token"))
	req := request.HeartbeatRequest{}
	json.Unmarshal(u.Ctx.Input.RequestBody, &req)

	cmd := heartbeat.Report(&req)

	rsp := response.HeartbeatResponse{}
	rsp.Code = 0
	rsp.Msg = "Success"
	rsp.Cmd = cmd

	u.Data["json"] = rsp
	u.ServeJSON()
}
