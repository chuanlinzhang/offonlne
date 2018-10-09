package routers

import (
	"bs/myfilms/controllers"
	"github.com/astaxie/beego"
	"bs/myfilms/controllers/customers"
	"bs/myfilms/controllers/films"
	"bs/myfilms/controllers/tickets"
	"bs/myfilms/controllers/orders"
	"bs/myfilms/controllers/admins"
)

func init() {
	// beego.Router("/", &controllers.MainController{})
	beego.Router("/", &controllers.BaseController{})

	beego.Router("/customers/login", &customers.LoginController{})
	beego.Router("/customers/login/check", &customers.LoginController{}, "Post:Login")
	beego.Router("/customers/exit", &customers.ExitController{}, "Get:Exit")
	beego.Router("/forgetpassword/into",&customers.ForgetPwdController{})
	beego.Router("/forgetpassword/get",&customers.ForgetPwdController{},"Post:GetAuth")
	beego.Router("/forgetpassword/auth",&customers.ForgetPwdController{},"Post:Auth")

	beego.Router("/customers/register", &customers.RegisterController{})
	beego.Router("/customers/register/check", &customers.RegisterController{}, "Post:Register")

	beego.Router("/customers-ranking1", &films.LookfilmController{}, "Get:Lookflim1")
	beego.Router("/customers-ranking2", &films.LookfilmController{}, "Get:Lookflim2")
	//beego.Router("/customers-ranking3", &films.LookfilmController{}, "Get:Lookflim3")

	beego.Router("/customers-function1", &customers.ChangeController{}, "Get:Change1")
	beego.Router("/customers-function2", &customers.ChangeController{}, "Get:Change2")
	beego.Router("/customers-function2/check", &customers.ChangeController{}, "Post:Change")
	beego.Router("/customers-function3", &customers.ChangeController{}, "Get:Change3")
	beego.Router("/customers-function3/check", &customers.ChangeController{}, "Post:ChangePwd")
	beego.Router("/customers-function4", &customers.ChangeController{}, "Get:Change4")
	beego.Router("/customers-function4/check", &customers.ChangeController{}, "Post:TopUp")

	beego.Router("/customers-tickets1", &tickets.ChangeTicketController{}, "Get:ChangeTicket1")
	beego.Router("/customers-tickets2", &tickets.ChangeTicketController{}, "Get:ChangeTicket2")
	beego.Router("/customers-tickets21", &tickets.ChangeTicketController{}, "Get:ChangeTicket2")
	beego.Router("/customers-tickets22", &tickets.ChangeTicketController{}, "Get:ChangeTicket22")
	beego.Router("/customers-tickets3", &tickets.ChangeTicketController{}, "Get:ChangeTicket3")

	beego.Router("/customers-order1", &orders.ChangeOrderController{}, "Get:ChangeOrder1")
	beego.Router("/customers-order2", &orders.ChangeOrderController{}, "Get:ChangeOrder1")
	beego.Router("/customers-order-look",&orders.ChangeOrderController{}, "Get:AddOrder")
	beego.Router("/customers-order-look1",&orders.ChangeOrderController{},"*:LookOrder")
	beego.Router("/customers-order-del",&orders.ChangeOrderController{}, "Get:DelOrder")
	beego.Router("/customers-order-grade",&orders.ChangeOrderController{},"Get:Grade")

	beego.Router("/admins/login", &admins.LoginController{})
	beego.Router("/admins/login/check", &admins.LoginController{}, "Post:Login")
	beego.Router("/admins/exit", &admins.ExitController{}, "Get:Exit")

	beego.Router("/admins/register", &admins.RegisterController{})
	beego.Router("/admins/register/check", &admins.RegisterController{}, "Post:Register")

	beego.Router("/admins-function1", &admins.ChangeController{}, "Get:Change1")
	beego.Router("/admins-function2", &admins.ChangeController{}, "Get:Change2")
	beego.Router("/admins-function2/check", &admins.ChangeController{}, "Post:Change")
	beego.Router("/admins-function3", &admins.ChangeController{}, "Get:Change3")
	beego.Router("/admins-function3/check", &admins.ChangeController{}, "Post:ChangePwd")

	beego.Router("/admins-movie1", &films.AdminfilmController{}, "Get:Adminflim1")
	beego.Router("/admins-movie2", &films.AdminfilmController{}, "Get:Adminflim2")
	beego.Router("/admins-movie2/check", &films.AdminfilmController{}, "Post:Addfilms")
	beego.Router("/admins-movie3", &films.AdminfilmController{}, "Get:Adminflim3")
	beego.Router("/admins-movie3/change", &films.AdminfilmController{}, "Get:Changefilms")
	beego.Router("/admins-movie3/change/check", &films.AdminfilmController{}, "Post:ChangeCheck")
	beego.Router("/admins-movie4", &films.AdminfilmController{}, "Get:Adminflim4")
	beego.Router("/admins-movie4/del", &films.AdminfilmController{}, "Get:Deletefilms")

	beego.Router("/admins-member1", &admins.ChangeController{}, "Get:ChangeCustomers1")
	beego.Router("/admins-member2", &admins.ChangeController{}, "Get:ChangeCustomers2")
	beego.Router("/admins-member2/delCus", &admins.ChangeController{}, "Get:DelCus")


}
