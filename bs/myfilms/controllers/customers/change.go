package customers

import (
	"bs/myfilms/models/users"
	"github.com/astaxie/beego/logs"
	"strconv"
	"crypto/md5"
	"fmt"
	"time"
	"strings"
)

type ChangeController struct {
	CustomersController
}

func (this *ChangeController) Get() {

}
func (this *ChangeController) Change() {
	cname := this.GetString("cname")
	cemail := this.GetString("cemail")
	nick_name := this.GetString("nick_name")
	name := this.GetString("name")
	mobile := this.GetString("mobile")
	sex := this.GetString("sex")
	b := users.Change(cname, cemail, nick_name, name, mobile, sex)
	if b == false {
		logs.Info("用户信息修改失败")
		this.TplName = "customers-function1.html"
		return
	}

	customers := users.Login(cname)
	this.SetSession("customers", customers)
	logs.Info("用户信息修改成功")
	this.Redirect("/customers-function1", 302)
}
func (this *ChangeController) ChangePwd() {
	cookie:=this.Ctx.Request.Header.Get("Cookie")
	cookieslice:=strings.Split(cookie,"; ")
	tk:=this.GetSession("token").(string)

			if cookieslice[0]==tk||cookieslice[1]==tk{
				logs.Info("token 验证成功")
			}else {
				logs.Info("token 验证失败")
				this.Redirect("/customers/login", 302)
				return
			}



	oldPwd := this.GetString("oldPwd")
	newPwd1 := this.GetString("newPwd1")
	newPwd2 := this.GetString("newPwd2")
	c := this.GetSession("customers").(*users.Customers) //这种可以
	if newPwd1 != newPwd2 {
		logs.Info("两次输入密码不同")
		this.Redirect("/customers-function3",302)
		return
	}
	if c.Pwd != oldPwd {
		logs.Info("原密码输入错误")
		this.Redirect("/customers-function3",302)
		return
	}
	b := users.ChangePwd(c.LoginName,c.Email,newPwd1)
	if b == false {
		logs.Info("用户密码修改失败")
		this.Redirect("/customers-function3",302)
		return
	}
	logs.Info("用户密码修改成功")
	this.DelSession("customers")
	this.Redirect("/customers/login", 302)
}
func (this *ChangeController) TopUp() {
	addBalance := this.GetString("addBalance")
	addB, err := strconv.ParseFloat(addBalance, 64)
	if err != nil {
		logs.Info("数据转换错误，充值失败")
		this.TplName = "customers-function4.html"
		return
	}
	pwd := this.GetString("pwd")
	c := this.GetSession("customers").(*users.Customers)
	sum := addB + c.Balance
	b := users.TopUp(c.LoginName, pwd, sum)
	if b==false{
		logs.Info("充值失败")
		this.TplName="customers-function4.html"
		return
	}
	customers := users.Login(c.LoginName)
	this.SetSession("customers", customers)
	logs.Info("用户充值成功")
	this.Redirect("/customers-function1", 302)

}
func (this *ChangeController) Change1() {

	customers := this.GetSession("customers")
	this.Data["customers"] = customers
	cu:=customers.(*users.Customers)
	data := []byte(cu.Name+cu.Pwd+time.Now().UTC().String())

	has := md5.Sum(data)
	token := fmt.Sprintf("%x", has)
	this.SetSession("token",token)
	this.Ctx.ResponseWriter.Header().Set("set-cookie",token)

	this.TplName = "customers-function1.html"
}
func (this *ChangeController) Change2() {
	customers := this.GetSession("customers")
	this.Data["customers"] = customers
	this.TplName = "customers-function2.html"
}
func (this *ChangeController) Change3() {
	customers := this.GetSession("customers")
	this.Data["customers"] = customers
	this.TplName = "customers-function3.html"
}
func (this *ChangeController) Change4() {
	customers := this.GetSession("customers")
	this.Data["customers"] = customers
	this.TplName = "customers-function4.html"
}
