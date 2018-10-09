package films

import (
	"bs/myfilms/models/films"

	"github.com/astaxie/beego/logs"
	"strconv"
)

type AdminfilmController struct {
	FilmsController
}

func (this *AdminfilmController) Get() {

}
func (this *AdminfilmController) Adminflim1() {
	list := make([]films.Films, 0)
	list = films.FilmsTreeGrid()
	this.Data["list"] = list
	admins := this.GetSession("admins")
	this.Data["admins"] = admins
	this.TplName = "admins-movie1.html"
}
func (this *AdminfilmController) Adminflim2() {
	n := films.Sumfilms()
	if n == -1 {
		logs.Info("访问数据条数失败")
		this.Redirect("/admins-movie1", 302)
		return
	}
	sum := n + 1
	this.Data["sum"] = sum
	admins := this.GetSession("admins")
	this.Data["admins"] = admins
	this.TplName = "admins-movie2.html"
}

//添加影片，及同步影票
func (this *AdminfilmController) Addfilms() {
	add(this)
}
func (this *AdminfilmController) Adminflim3() {

	admins := this.GetSession("admins")
	this.Data["admins"] = admins
	this.Redirect("/admins-movie1", 302)
}
func (this *AdminfilmController) Changefilms() {
	FilmsNo := this.GetString("FilmsNo")
	f := films.Changefilms(FilmsNo)
	if f == nil {
		logs.Info("进入修改失败")
		this.Redirect("/admins-movie1", 302)
		return
	}
	logs.Info("进入修改")
	this.Data["films"] = f
	this.TplName = "admins-movie3.html"
}
func (this *AdminfilmController) ChangeCheck() {
	filmsNo := this.GetString("filmsNo")
	b := films.Deletefilms(filmsNo)
	f,_:=strconv.Atoi(filmsNo)
	films.DeleteTickets(f)
	if b == true {
		add(this)
	}

}
func add(this *AdminfilmController) {
	//filmsNo := this.GetString("filmsNo")
	//nameCh := this.GetString("nameCh")
	//nameEn := this.GetString("nameEn")
	//director := this.GetString("director")
	//actor := this.GetString("actor")
	//filmLength := this.GetString("filmLength")
	//Pricefull := this.GetString("Pricefull")
	//Pricest := this.GetString("Pricest")
	//Intro := this.GetString("Intro")
	//Hot:=this.GetString("hot")
	//New :=this.GetString("new")
	obj := films.Films{}
	this.ParseForm(&obj)
	b := films.Addfilms1(obj)
	go films.MakeTickets(obj)
	//filmLength1, err := strconv.Atoi(filmLength)
	//if err != nil {
	//	logs.Info("数据类型转化错误，添加影片失败")
	//	this.Redirect("/admins-movie2", 302)
	//	return
	//}
	//Hot1, err := strconv.Atoi(Hot)
	//if err != nil {
	//	logs.Info("数据类型转化错误，添加影片失败")
	//	this.Redirect("/admins-movie2", 302)
	//	return
	//}
	//New1, err := strconv.Atoi(New)
	//if err != nil {
	//	logs.Info("数据类型转化错误，添加影片失败")
	//	this.Redirect("/admins-movie2", 302)
	//	return
	//}
	//Pricefull1, err := strconv.ParseFloat(Pricefull, 64)
	//if err != nil {
	//	logs.Info("数据类型转化错误，添加影片失败")
	//	this.Redirect("/admins-movie2", 302)
	//	return
	//}
	//Pricest1, err := strconv.ParseFloat(Pricest, 64)
	//if err != nil {
	//	logs.Info("数据类型转化错误，添加影片失败")
	//	this.Redirect("/admins-movie2", 302)
	//	return
	//}
	//b := films.Addfilms(filmsNo, nameCh, nameEn, director, actor, Intro, filmLength1, Hot1,New1,Pricefull1, Pricest1)
	if b == false {
		logs.Info("添加影片失败")
		this.Redirect("/admins-movie2", 302)
		return
	}
	logs.Info("影片添加成功")
	this.Redirect("/admins-movie1", 302)
}
func (this *AdminfilmController) Adminflim4() {
	admins := this.GetSession("admins")
	this.Data["admins"] = admins
	this.Redirect("/admins-movie1", 302)
}
func (this *AdminfilmController) Deletefilms() {
	FilmsNo := this.GetString("FilmsNo")
	b := films.Deletefilms(FilmsNo)
	if b == false {
		logs.Info("删除失败")
		this.Redirect("/admins-movie1", 302)
		return
	}
	logs.Info("删除成功")
	this.Redirect("/admins-movie1", 302)
	return
}
