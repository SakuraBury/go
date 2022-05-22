package controllers

import (
	"common/models"
	"encoding/json"
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	"open-center/common"
	"open-center/request"
)

type JobController struct {
	beego.Controller
}

func (u *JobController) AddJob() {
	rsp := models.BaseResult{}
	rsp.Code = 0
	rsp.Msg = "Success"

	req := request.AddJobRequest{}
	err := json.Unmarshal(u.Ctx.Input.RequestBody, &req)
	if err != nil {
		fmt.Println(err)
		rsp.Code = -1
		rsp.Msg = err.Error()
		u.Data["json"] = rsp
		u.ServeJSON()

		return
	}

	producer := common.GetMockProducerInstance()
	cmd := producer.GenerateCmd(req.ProcessorId, req.Endpoint, req.Target, req.Params)
	producer.PushCmd(cmd)

	rsp.Data = cmd

	u.Data["json"] = rsp
	u.ServeJSON()
}
