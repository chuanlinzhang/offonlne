package routers

import (
	"eoin/userservice/controllers"
	"github.com/astaxie/beego"
)

func init() {
	//进入首页
	beego.Router("/", &controllers.MainController{})


    //进入登录页面
    beego.Router("/user/login",&controllers.LoginController{})
    //登录验证
	beego.Router("/user/login/check",&controllers.LoginController{},"Post:Login")


    //进入修改密码页面
	beego.Router("/user/change",&controllers.ChangController{})
    //修改验证，通过并保存
	beego.Router("/user/change",&controllers.ChangController{},"Post:ChangPWD")


	//进入注册页面
	beego.Router("/user/register",&controllers.RegisterController{})
	//注册检查
	beego.Router("/user/register/check",&controllers.RegisterController{},"Post:Registe")
}
