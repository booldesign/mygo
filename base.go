package controllers

import (
	"github.com/astaxie/beego"
	// "strconv"
	"strings"
	// "time"
)

type BaseController struct {
	beego.Controller
}

func (this *BaseController) Finish() {

}

func (this *BaseController) Prepore() {

}

//是否post提交
func (this *BaseController) isPost() bool {
	return this.Ctx.Request.Method == "POST"
}

//获取用户IP地址
func (this *BaseController) getClientIp() string {
	s := strings.Split(this.Ctx.Request.RemoteAddr, ":")

	return s[0]
}

//权限验证
// func (this *BaseController) checkPermission() {
// 	if this.userid != 1 && this.controllerName == "user" {
// 		this.showmsg("抱歉，只有超级管理员才能进行该操作！")
// 	}
// }

// func (this *BaseController) getTime() time.Time {
// 	timezone, _ := strconv.ParseFloat(option.Get("timezone"), 64)
// 	add := timezone * float64(time.Hour)
// 	return time.Now().UTC().Add(time.Duration(add))
// }
