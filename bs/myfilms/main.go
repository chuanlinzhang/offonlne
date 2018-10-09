package main

import (
	_ "bs/myfilms/routers"

	"github.com/beego/i18n"
	"github.com/astaxie/beego"
)

func main() {
	beego.BConfig.WebConfig.Session.SessionOn=true
	//注册模板功能
	beego.AddFuncMap("i18n", i18n.Tr)
	beego.Run()
}

