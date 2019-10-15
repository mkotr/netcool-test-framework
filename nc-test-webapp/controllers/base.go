package controllers

import (
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego"
)

const (
	SUCCESS       = 0
	UNKNOWN       = 1
	PARSEERROR    = 98
	SESSIONEXPIRE = 99
)

type ResResult struct {
	Status  int         `json:"status"`
	Data    interface{} `json:"data"`
	Message string      `json:"msg"`
}

type BaseController struct {
	beego.Controller
}

func (b *BaseController) ParseBody(v interface{}) error {
	fmt.Println(string(b.Ctx.Input.RequestBody))
	return json.Unmarshal(b.Ctx.Input.RequestBody, v)
}

func (b *BaseController) ResParseError(err error) {
	b.Error(PARSEERROR, "The requested data format is incorrect!", err)
}

func (b *BaseController) ResJson(v interface{}) {
	b.Data["json"] = v
	b.ServeJSON()
}

func (b *BaseController) Success(data interface{}, msg string) {
	result := ResResult{
		Status:  SUCCESS,
		Data:    data,
		Message: msg,
	}

	b.ResJson(result)
}

//SuccessMultiple - when you want to return more than one data obj..
func (b *BaseController) SuccessMultiple(data map[string]interface{}, msg string) {
	result := ResResult{
		Status:  SUCCESS,
		Data:    data,
		Message: msg,
	}

	b.ResJson(result)
}

func (b *BaseController) Error(status int, msg string, err error) {
	result := ResResult{
		Status:  status,
		Message: msg,
	}

	beego.Error(fmt.Sprintf("%s:%v", msg, err))

	b.ResJson(result)
}
