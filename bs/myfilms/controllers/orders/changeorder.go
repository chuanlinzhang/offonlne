package orders

import (
	"bs/myfilms/models/films"
	"bs/myfilms/models/users"
	"strconv"
	"github.com/astaxie/beego/logs"
	"bs/myfilms/models/sendemail"
)

type ChangeOrderController struct {
	OrdersController
}

func (this *ChangeOrderController) Get() {

}
func (this *ChangeOrderController) ChangeOrder() {

}
func (this *ChangeOrderController) ChangeOrder1() {
	customers := this.GetSession("customers")
	this.Data["customers"] = customers
	this.TplName = "customers-order1.html"
}

//func (this *ChangeOrderController) ChangeOrder2()  {
//	customers:=this.GetSession("customers")
//	this.Data["customers"]=customers
//	this.TplName="customers-order1.html"
//}
func (this *ChangeOrderController) AddOrder() {
	customers := this.GetSession("customers").(*users.Customers)
	Tno:=this.GetString("Tno")
	NameCh := this.GetString("NameCh")
	Amount := this.GetString("Amount")
	Amount1, _ := strconv.ParseFloat(Amount, 64)
	StartTime := this.GetString("StartTime")
	Seat := this.GetString("Seat")
	b := films.AddOrders(customers, NameCh, StartTime, Seat, Amount1)

	if b == false {
		logs.Info("添加订单失败")
		this.Redirect("/customers-tickets2", 302)
		return
	}
	go sendemail.BuyTicket(customers.LoginName,customers.Email)
	logs.Info("添加订单成功")
	t,_:=strconv.Atoi(Tno)
	_=films.DeleteTickets(t)
	this.LookOrder()
}
func (this *ChangeOrderController) LookOrder() {
	customers := this.GetSession("customers").(*users.Customers)
	this.Data["customers"]=customers
	orders := films.LookOrders(customers.LoginName)
	this.Data["orders"] = orders
	this.TplName = "customers-order1.html"
}
func (this *ChangeOrderController) DelOrder() {
	OrderNo := this.GetString("OrderNo")
	o, _ := strconv.Atoi(OrderNo)
	b := films.DelOrders(o)
	if b == false {
		logs.Info("删除订单失败")
		this.Redirect("/customers-tickets1", 302)
		return
	}
	logs.Info("删除订单成功")
	this.Redirect("/customers-order-look1", 302)
}
func (this *ChangeOrderController) Grade() {
	customers := this.GetSession("customers").(*users.Customers)
	this.Data["customers"]=customers
	OrderNo:=this.GetString("OrderNo")
	this.Data["OrderNo"]=OrderNo
	this.TplName="customers-tickets3.html"
}