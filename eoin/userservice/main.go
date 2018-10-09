package main

import (
	_ "eoin/userservice/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"eoin/userservice/models"
	"github.com/beego/i18n"
)

func init() {
	//注册一个引擎
	logs.Register("mongo", func() logs.Logger {
		return &models.MongoDBLogger{}
	})
	//设置.后面的在调用日志的相关方法会把信息写入到数据库
	beego.SetLogger("mongo",`{"db_name":"llog"}`)
}
func main() {
	beego.BConfig.WebConfig.Session.SessionOn = true
	//用beego.Info打印日志到文件
	beego.Info(beego.BConfig.AppName, "0.1.1.0227")
	//注册模板功能
	beego.AddFuncMap("i18n", i18n.Tr)
	beego.Run()
}

