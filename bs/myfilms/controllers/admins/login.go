package admins

import (

	"bs/myfilms/models/users"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego"

)

type LoginController struct {
	AdminController
}

func (this *LoginController) Get() {
	this.TplName = "admin-login.html"
}
func (this *LoginController) Login() {
	admin_login := this.GetString("admin_login")
	admin_pwd := this.GetString("admin_pwd")

	if admin_login == "" && admin_pwd == "" {
		this.Redirect(beego.URLFor("admins.LoginController.Get"), 302)
		return
	}
	a := users.LoginA(admin_login)
	if a == nil {
		logs.Info("用户不存在")
		this.Redirect(beego.URLFor("admins.LoginController.Get"), 302)
		return
	}
	if a.AdminPwd != admin_pwd {
		logs.Info("账号或者密码错误")
		this.Redirect(beego.URLFor("admins.LoginController.Get"), 302)
		return
	}
	this.SetSession("admins",a)
	//this.TplName="admins-function1.html"
	this.Redirect(beego.URLFor("admins.ChangeController.Change1"),302)

}
