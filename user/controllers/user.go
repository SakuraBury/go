package controllers

import (
	models "common/models"
	"encoding/json"
	"strconv"
	"user/dal/model"

	beego "github.com/beego/beego/v2/server/web"
)

// Operations about Users
type UserController struct {
	beego.Controller
}

// @Title CreateUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router / [post]
func (u *UserController) Add() {
	var user model.TUserDTO
	json.Unmarshal(u.Ctx.Input.RequestBody, &user)
	err := user.CreateUser()

	result := models.BaseResult{}
	if err != nil {
		result.Code = -1
		result.Msg = "插入用户错误"
	} else {
		result.Code = 0
		result.Msg = "Success"
		result.Data = user.ID
	}

	u.Data["json"] = result
	u.ServeJSON()
}

// @Title Get
// @Description get user by uid
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :uid is empty
// @router /:uid [get]
func (u *UserController) Get() {
	uid := u.GetString(":uid")

	var user = model.TUserDTO{}

	if uid != "" {
		id, _ := strconv.Atoi(uid)
		user.GetUserById(uint(id))
		u.Data["json"] = user
	}
	u.ServeJSON()
}

// @Title Login
// @Description Logs user into the system
// @Param	username		query 	string	true		"The username for login"
// @Param	password		query 	string	true		"The password for login"
// @Success 200 {string} login success
// @Failure 403 user not exist
// @router /login [get]
func (u *UserController) Login() {
	phone := u.GetString("phone")
	password := u.GetString("password")
	var user = model.TUserDTO{}

	user.GetUserByPhoneAndPassword(phone, password)

	if user.ID > 0 {
		u.Data["json"] = "login success"
	} else {
		u.Data["json"] = "user not exist"
	}
	u.ServeJSON()
}
