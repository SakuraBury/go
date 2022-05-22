package filter

import (
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
)

func AuthFilter(ctx *context.Context) {
	if ctx.Request.RequestURI == "user/v1/login" {
		return
	}

	token := ctx.Request.Header.Get("auth-token")
	if token == "" {
		logs.Error("====AuthFilter no auth token.uri:%s", ctx.Request.RequestURI)
		ctx.Redirect(401, "miss auth token")
		return
	}

	if "asdfghjkl" != token {
		logs.Error("====AuthFilter  auth token err: %s, uri:%s", token, ctx.Request.RequestURI)
		ctx.Redirect(401, "error auth token")
		return
	}

	ctx.Input.SetData("token", token)
}

func init() {
	beego.InsertFilter("*", beego.BeforeRouter, AuthFilter)
}
