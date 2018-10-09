package controllers

type LoginController struct {
	baseController
}

func (this *LoginController) Get()  {
	this.TplName= "login.html"

}
func (this *LoginController) Login()  {


}