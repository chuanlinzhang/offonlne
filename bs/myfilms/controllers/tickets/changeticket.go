package tickets

import (
	"bs/myfilms/models/films"

)

type ChangeTicketController struct {
	TicketsController
}

func (this *ChangeTicketController) Get() {

}
func (this *ChangeTicketController) ChangeTicket() {

}
func (this *ChangeTicketController) ChangeTicket1() {
	customers:=this.GetSession("customers")
	this.Data["customers"]=customers
	this.TplName = "customers-tickets1.html"
}
func (this *ChangeTicketController) ChangeTicket2() {
	FilmsNo:=this.GetString("FilmsNo")
   tickets:=films.LookTickets1(FilmsNo)
    this.Data["FilmsNo"]=FilmsNo
   this.Data["tickets"]=tickets
	customers:=this.GetSession("customers")
	this.Data["customers"]=customers
	this.TplName = "customers-tickets2.html"
}
func (this *ChangeTicketController) ChangeTicket22() {
	FilmsNo:=this.GetString("FilmsNo")
	tickets:=films.LookTickets2(FilmsNo)
	this.Data["FilmsNo"]=FilmsNo
	this.Data["tickets"]=tickets
	customers:=this.GetSession("customers")
	this.Data["customers"]=customers
	this.TplName = "customers-tickets2.html"
}
func (this *ChangeTicketController) ChangeTicket3() {
	customers:=this.GetSession("customers")
	this.Data["customers"]=customers
	this.TplName = "customers-tickets3.html"
}
