package admins

import (
	"github.com/astaxie/beego"
	"bs/myfilms/models/users"

)

type RegisterController struct {
	AdminController
}

func (this *RegisterController) Get() {
	this.TplName = "admin-register.html"
}
func (this *RegisterController) Register() {
	//admin_login := this.GetString("admin_login")
	admin_pwd1 := this.GetString("admin_pwd1")
	admin_pwd2 := this.GetString("admin_pwd2")
	//admin_email := this.GetString("admin_email")
	//admin_nickname := this.GetString("admin_nickname")
	//admin_name := this.GetString("admin_name")
	//admin_mobile := this.GetString("admin_mobile")
	//admin_sex := this.GetString("admin_sex")
	if admin_pwd1 != admin_pwd2 {
		this.Redirect(beego.URLFor("/admins/register"), 302)
		return
	}
	admin:=users.Admins{}
      this.ParseForm(&admin)


	//b := users.RegisterA(admin_login, admin_pwd1, admin_nickname, admin_name, admin_mobile, admin_sex, admin_email)
	b:=users.RegisterA1(admin)
	if b == false {
		this.Redirect(beego.URLFor("/admins/register"), 302)
		return
	}
	this.TplName = "admin-login.html"
}
