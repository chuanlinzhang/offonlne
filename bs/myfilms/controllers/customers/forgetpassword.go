package customers

import (

	"github.com/astaxie/beego/logs"
	"bs/myfilms/models/users"
	"time"
	"fmt"
	"math/rand"
	"bs/myfilms/models/sendemail"
)

type ForgetPwdController struct {
	CustomersController
}


var authCode string
var customer *users.Customers
func (f *ForgetPwdController) Get()  {
	f.TplName="fotgetpassword.html"
}
func (f *ForgetPwdController) GetAuth()  {
	login_name:=f.GetString("login_name")
	if login_name==""{
		logs.Info("输入账户为空")
		return
	}
	customer=users.Login(login_name)
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	authCode= fmt.Sprintf("%06v", rnd.Int31n(1000000))
	go sendemail.GetAuth(customer.LoginName,customer.Email,authCode)
   f.TplName="forgetpasswordemail.html"
}
func (f *ForgetPwdController) Auth()  {
	auth:=f.GetString("auth")
	if auth==authCode{
		logs.Info("验证码正确")
		go sendemail.CheckAuth(customer.LoginName,customer.Email,customer.Pwd)
		f.Redirect("/customers/login",302)
		return
	}else {
		logs.Info("验证码错误")
		f.Redirect("/forgetpassword/get",302)
		return
	}

}