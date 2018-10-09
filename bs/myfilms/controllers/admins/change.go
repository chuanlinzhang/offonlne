package admins

import (
	"bs/myfilms/models/users"
	"github.com/astaxie/beego/logs"

)

type ChangeController struct {
	AdminController
}

func (thhis *ChangeController) Get() {

}
func (this *ChangeController) Change() {
	//adminLogin := this.GetString("adminLogin")
	//adminNickName := this.GetString("adminNickName")
	//adminName := this.GetString("adminName")
	//adminMobile := this.GetString("adminMobile")
	//adminSex := this.GetString("adminSex")
	//adminEmail := this.GetString("adminEmail")
	//admins:=this.GetSession("admins").(users.Admins)
	admin:=users.Admins{}
	this.ParseForm(&admin)
	b := users.ChangeA(admin.AdminLogin, admin.AdminNickName, admin.AdminName, admin.AdminMobile, admin.AdminSex, admin.AdminEmail)
	if b == false {
		logs.Info("管理员信息修改失败")
		this.Redirect("/admins-function1", 302)
		return
	}
	admins := users.LoginA(admin.AdminLogin)
	this.SetSession("admins", admins)
	logs.Info("管理员信息修改成功")
	this.Redirect("/admins-function1", 302)
}
func (this *ChangeController) ChangePwd() {
	oldPwd := this.GetString("oldPwd")
	newPwd1 := this.GetString("newPwd1")
	newPwd2 := this.GetString("newPwd2")
	a := this.GetSession("admins").(*users.Admins)
	if newPwd1 != newPwd2 {
		logs.Info("两次密码输入不一样")
		this.Redirect("/admins-function3", 302)
		return
	}
	if a.AdminPwd != oldPwd {
		logs.Info("原密码输入错误")
		this.Redirect("/admins-function3", 302)
		return
	}
	b := users.ChangePwdA(a.AdminLogin, newPwd1)
	if b == false {
		logs.Info("密码修改失败")
		this.Redirect("/admins-function3", 302)
		return
	}
	logs.Info("密码修改成功")
	this.DelSession("admins")
	this.Redirect("/admins/login", 302)
}
func (this *ChangeController) Change1() {
	admins := this.GetSession("admins")
	this.Data["admins"] = admins
	this.TplName = "admins-function1.html"
}
func (this *ChangeController) Change2() {
	admins := this.GetSession("admins")
	this.Data["admins"] = admins
	this.TplName = "admins-function2.html"
}
func (this *ChangeController) Change3() {
	admins := this.GetSession("admins")
	this.Data["admins"] = admins
	this.TplName = "admins-function3.html"
}

func (this *ChangeController) ChangeCustomers1() {
	//var result [] *users.Customers
	list := make([]users.Customers, 0)
	list = users.CustomersTreeGrid()
	this.Data["list"] = list
	admins := this.GetSession("admins")
	this.Data["admins"] = admins
	this.TplName = "admins-member1.html"
}
func (this *ChangeController) ChangeCustomers2() {
	list := make([]users.Customers, 0)
	list = users.CustomersTreeGrid()
	this.Data["list"] = list
	admins := this.GetSession("admins")
	this.Data["admins"] = admins
	this.TplName = "admins-member2.html"
}
func (this *ChangeController) DelCus() {
	LoginName:=this.GetString("LoginName")
	b:=users.DelCus(LoginName)
	if b==false{
		logs.Info("删除失败")
		this.Redirect("/admins-member1",302)
		return
	}
	logs.Info("删除成功")
	list := make([]users.Customers, 0)
	list = users.CustomersTreeGrid()
	this.Data["list"] = list
	admins := this.GetSession("admins")
	this.Data["admins"] = admins
	this.TplName = "admins-member2.html"

}
