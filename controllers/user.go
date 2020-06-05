package controllers

import (
	"github.com/astaxie/beego/logs"
)

type UserController struct {
	BaseController
}

func (ctrl *UserController) Login() {
	username := ctrl.GetString("username")
	password := ctrl.GetString("password")

	logs.Debug("user is %s password is %s", username, password)

	testParam := ctrl.Ctx.Input.GetData("test_param")
	logs.Debug("test param is %v", testParam)

	ctrl.apiSuccess(map[string]interface{}{
		"msg": "SUCCESS",
		"token": "xxxxxxxxxxxxxxxxxxxxxx" ,
	})
}

func (ctrl *UserController) Test() {


	ctrl.apiSuccess(map[string]interface{}{
		"msg": "SUCCESS",
	})
}