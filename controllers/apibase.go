package controllers

import (
	"github.com/astaxie/beego"
)

type ApibaseController struct {
	beego.Controller
}

func (this *ApibaseController) Prepare() {
	beego.LoadAppConfig("ini", "conf/web.conf")
	this.Data["Website"] = beego.AppConfig.String("Website")
}
