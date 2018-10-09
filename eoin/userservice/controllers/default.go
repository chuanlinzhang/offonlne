package controllers

import (
	"github.com/astaxie/beego"
	"strings"
	"github.com/beego/i18n"
	"fmt"
)

var langTypes []string

type MainController struct {
	baseController
}

func init() {
	langTypes = strings.Split(beego.AppConfig.String("lang_types"), "|")
	for _, lang := range langTypes {
		beego.Trace("Loading language: " + lang)
		if err := i18n.SetMessage(lang, "conf/"+"locale_"+lang+".ini"); err != nil {
			beego.Error("Fail to set message file:", err)
			return
		}
	}
}

type baseController struct {
	beego.Controller
	i18n.Locale
}

func (this *baseController) Prepare() {
	this.Lang = ""
	al := this.Ctx.Request.Header.Get("Accept-Language")
	if len(al) > 4 {
		al = al[:5]
		if i18n.IsExist(al) {
			this.Lang = al
		}
	}
	if len(this.Lang) == 0 {
		this.Lang = "en-US"
	}
	this.Data["Lang"] = this.Lang
	fmt.Println(this.Lang,"*****************")
}

func (this *MainController) Get() {
	//c.Data["Website"] = "beego.me"
	//c.Data["Email"] = "astaxie@gmail.com"
	this.TplName = "welcome.html"
}
