package customers

type ExitController struct {
	CustomersController
}

func (this *ExitController) Exit() {
	this.DelSession("customers")
	this.TplName="login-index.html"
}
