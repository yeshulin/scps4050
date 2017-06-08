package main

import (
	"webproject/4050/common/function"

	_ "webproject/4050/initial"

	_ "webproject/4050/routers"

	"github.com/astaxie/beego"
)

func main() {
	beego.AddFuncMap("ConvertT", function.ConvertT)
	beego.SetStaticPath("/uploads", "uploads")
	//	beego.InsertFilter("/", beego.BeforeRouter, function.TransparentStatic)
	//	beego.InsertFilter("/*", beego.BeforeRouter, function.TransparentStatic)

	beego.Run()
}
