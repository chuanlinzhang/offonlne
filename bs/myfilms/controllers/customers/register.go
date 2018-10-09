package customers

import (
	"github.com/astaxie/beego"
	"bs/myfilms/models/users"

	"fmt"
	"github.com/astaxie/beego/validation"
	"log"
)

type RegisterController struct {
	CustomersController
}

func (this *RegisterController) Get() {
	this.TplName = "customers-register.html"
}
func (this *RegisterController) Register() {

	//cname := this.GetString("cname")
	cpwd1 := this.GetString("cpwd1")
	cpwd2 := this.GetString("cpwd2")
	//cemail := this.GetString("cemail")
	//nick_name := this.GetString("nick_name")
	//name := this.GetString("name")
	//mobile := this.GetString("mobile")
	//sex := this.GetString("sex")

	if cpwd1 != cpwd2 {
		log.Println("密码不一致")
		this.Redirect(beego.URLFor("customers.RegisterController.Get"), 302)
		return
	}
	customers:=users.Customers{}
	err:=this.ParseForm(&customers)
	if err!=nil{
		log.Println("解析表单出错")
		this.Redirect(beego.URLFor("customers.RegisterController.Get"), 302)
		return
	}
	valid:=validation.Validation{}
	b,err:=valid.Valid(&customers)
	if err==nil&&b{
		b:=users.Register1(customers)
		if b==false{
			fmt.Println("b",b)
			this.Redirect(beego.URLFor("customers.RegisterController.Get"),302)
			return
		}
		this.TplName="login-index.html"
	}else{
		log.Println("邮箱或者电话号码格式不符合")
		this.Redirect(beego.URLFor("customers.RegisterController.Get"), 302)
		return
	}
  // b:=users.Register(cname,cpwd1,cemail,nick_name,name,mobile,sex)


}
