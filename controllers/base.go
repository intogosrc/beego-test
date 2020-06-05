package controllers

import (
	"github.com/astaxie/beego"
)

const (
	API_SUCCESS = 0
	API_ERR_MSG = 500 // 显示错误信息
)

type ApiRet struct {
	Code int32                  `json:"code"`
	Data map[string]interface{} `json:"data"`
	Msg  string                 `json:"msg"`
}

type BaseController struct {
	beego.Controller
}

func (ctrl *BaseController) apiSuccess(data map[string]interface{}) {
	ctrl.Data["json"] = &ApiRet{
		API_SUCCESS,
		data,
		"",
	}
	ctrl.ServeJSON()
}

func (ctrl *BaseController) apiError(code int32, msg string) {
	ctrl.Data["json"] = &ApiRet{
		code,
		nil,
		msg,
	}
	ctrl.ServeJSON()
}