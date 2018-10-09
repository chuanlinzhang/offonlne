package films

import (
	"bs/myfilms/models/films"
	"github.com/astaxie/beego/logs"
)

type LookfilmController struct {
	FilmsController
}

func (this *LookfilmController) Get() {

}
func (this *LookfilmController) Lookflim1() {
	Hotfilms := make([]films.Films, 0)
	Hotfilms = films.Hotfilms()
	if Hotfilms == nil {
		logs.Info("热度查询失败")
		this.Redirect("/customers-function1", 302)
		return
	}
	this.Data["Hotfilms"] = Hotfilms
	customers := this.GetSession("customers")
	this.Data["customers"] = customers
	this.TplName = "customers-ranking1.html"
}
func (this *LookfilmController) Lookflim2() {
	newfilms := make([]films.Films, 0)
	newfilms=films.Newfilms()
	if newfilms==nil{
		logs.Info("新片查询失败")
		this.Redirect("/customers-function1",302)
		return
	}
	this.Data["newfilms"]=newfilms
	customers := this.GetSession("customers")
	this.Data["customers"] = customers
	this.TplName = "customers-ranking2.html"
}

//func (this *LookfilmController) Lookflim3()  {
//	customers:=this.GetSession("customers")
//	this.Data["customers"]=customers
//	this.TplName="customers-ranking3.html"
//}
