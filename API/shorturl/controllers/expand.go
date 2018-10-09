package controllers

import "github.com/astaxie/beego"

type ExpandController struct {
	beego.Controller
}

func (e *ExpandController) Get()  {
    var result ShortResult
    shorturl:=e.Input().Get("shorturl")
    result.Urlshort=shorturl
    if urlcache.IsExist(shorturl){
    	result.Urllong=urlcache.Get(shorturl).(string)
	}else {
		result.Urllong=""
	}
	e.Data["json"]=result
	e.ServeJSON()
}
