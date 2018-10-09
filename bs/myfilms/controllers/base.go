package controllers

import (
	"github.com/astaxie/beego"
	"strings"
	"github.com/beego/i18n"
)

type BaseController struct {
	beego.Controller
}

func Lang()  {
	langTypes := strings.Split(beego.AppConfig.String("lang_types"), "|")
	for _, lang := range langTypes {
		//beego.Trace("Loading language: " + lang)
		if err := i18n.SetMessage(lang, "conf/"+"locale_"+lang+".ini"); err != nil {
			beego.Error("Fail to set message file:", err)
			return
		}
	}
}
func (this *BaseController) Get()  {
	//this.TplName="404.html"
	this.Redirect(beego.URLFor("customers.LoginController.Get"),302)
}