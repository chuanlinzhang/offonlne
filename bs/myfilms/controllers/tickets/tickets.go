package tickets

import (
	"github.com/beego/i18n"
	"bs/myfilms/controllers"
)


func init() {
	controllers.Lang()
}

type TicketsController struct {
	controllers.BaseController
	i18n.Locale
}

//Prepare  Finish 相当于构造与析构函数
func (this *TicketsController) Prepare() {
	this.Lang = ""
	al := this.Ctx.Request.Header.Get("Accept-Language") //获取请求头
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

}
func (this *TicketsController) Finish()  {

}

