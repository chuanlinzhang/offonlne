package customers

import (

	"github.com/astaxie/beego"
	"bs/myfilms/models/users"
	"github.com/astaxie/beego/logs"
)

type LoginController struct {
	CustomersController
}

func (this *LoginController) Get() {

	this.TplName = "login-index.html"

}
func (this *LoginController) Login() {
	login_name := this.GetString("login_name")
	pwd := this.GetString("pwd")
	if login_name == "" && pwd == "" {
		this.Redirect(beego.URLFor("customers.LoginController.Get"), 302)
		return
	}
	customers := users.Login(login_name)
	if customers == nil {
		logs.Info("用户不存在")
		this.Redirect(beego.URLFor("customers.LoginController.Get"), 302)
		return
	}
	if customers.Pwd != pwd {
		logs.Info("账号或者密码错误")
		this.Redirect(beego.URLFor("customers.LoginController.Get"), 302)
		return
	}
	this.SetSession("customers", customers)
   //this.TplName="customers-function1.html"
   this.Redirect(beego.URLFor("customers.ChangeController.Change1"),302)
}
