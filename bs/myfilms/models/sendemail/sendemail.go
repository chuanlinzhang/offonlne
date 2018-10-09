package sendemail
import (
	"github.com/go-gomail/gomail"
	"log"
)

func sendEmail(cname,toEmail,subject,body string) bool {
	m:=gomail.NewMessage()
	m.SetAddressHeader("From","1558030788@qq.com","影票系统官方邮箱")
	m.SetHeader("To",
		m.FormatAddress(toEmail,cname))//接收地址与用户名
	m.SetHeader("Subject",subject)//主题
	m.SetBody("text/html",body)
	d:=gomail.NewDialer("smtp.qq.com",465,"1558030788@qq.com","zjvuwigwunurhfad")
	if err:=d.DialAndSend(m);err!=nil{
		return false
	}
	return true
}
//修改信息发生邮箱
func ChangeInfoSE(cname,toEmail string)  {
	subject:="用户基本信息修改"
	body:="用户基本信息已经修改成功，是否为本人操作，如果不是本人操作请点击"+"<a href = \"http://localhost:8080\">这里</a>"+"进行密码重置"
	b:=sendEmail(cname,toEmail,subject,body)
	if b==true{
		log.Println("邮件发生成功")
	}
}
//修改密码发生邮箱
func ChangePwdSE(cname,toEmail string){
	subject:="用户密码修改"
	body:="用户密码已经修改成功，是否为本人操作，如果不是本人操作请点击"+"<a href = \"http://localhost:8080\">这里</a>"+"找回密码"
	b:=sendEmail(cname,toEmail,subject,body)
	if b==true{
		log.Println("邮件发生成功")
	}
}

//购票成功发生邮箱
func BuyTicket(cname,toEmail string)  {
	subject:="用户购票"
	body:="用户购票成功，请点击"+"<a href = \"http://localhost:8080\">这里</a>"+"查看订单信息"
	b:=sendEmail(cname,toEmail,subject,body)
	if b==true{
		log.Println("邮件发生成功")
	}
}
//忘记密码获取验证码
func GetAuth(cname,toEmail,authCode string)  {
	subject:="获取随机验证码"
	body:="你好"+cname+"：本次随机验证码为："+authCode+"，10分钟有效"
	b:=sendEmail(cname,toEmail,subject,body)
	if b==true{
		log.Println("邮件发生成功")
	}
}
//忘记密码验证验证码
func CheckAuth(cname,toEmail,pwd string)  {
	subject:="密码"
	body:="你好"+cname+"：你的密码为："+pwd+"，请记住你的密码！"
	b:=sendEmail(cname,toEmail,subject,body)
	if b==true{
		log.Println("邮件发生成功")
	}
}