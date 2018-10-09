package controllers

import "fmt"

type RegisterController struct {
	baseController
}

func (this *RegisterController) Get()  {
	this.TplName="register.html"
}
func (this *RegisterController) Registe()  {
	username:=this.GetString("username")
	userpassword1:=this.GetString("userpassword1")
	userpassword2:=this.GetString("userpassword2")

fmt.Println(username,userpassword1,userpassword2)
}