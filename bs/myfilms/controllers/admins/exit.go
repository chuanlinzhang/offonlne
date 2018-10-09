package admins


type ExitController struct {
	AdminController
}

func (this *ExitController) Exit()  {
	this.DelSession("admins")
	this.TplName="login-index.html"
}