package controllers

import (
	"github.com/astaxie/beego/cache"
	"github.com/astaxie/beego"
	"API/shorturl/models"
)

var urlcache cache.Cache

func init()  {
	urlcache,_=cache.NewCache("memory",`{"interval":0}`)
}

type ShortResult struct {
	Urlshort string
	Urllong string
}
type ShortController struct {
	beego.Controller
}

func (s *ShortController) Get()  {
	var result ShortResult
	longurl:=s.Input().Get("longurl")
	beego.Info(longurl)
	result.Urllong=longurl
	urlmd5:=models.GetMD5("longurl")
	beego.Info(urlmd5)
	if urlcache.IsExist(urlmd5){
		result.Urlshort=urlcache.Get(urlmd5).(string)
	}else {
		result.Urlshort=models.Generate()
		err:=urlcache.Put(urlmd5,result.Urlshort,0)
		if err!=nil{
			beego.Info(err)
			}
			err=urlcache.Put(result.Urlshort,longurl,0)
			if err!=nil{
				beego.Info(err)
			}
	}
	s.Data["json"]=result
	s.ServeJSON()
}