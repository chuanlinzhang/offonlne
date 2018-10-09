package main
import (
	"net/smtp"
	"strings"

	"log"
	"fmt"
)


//发送邮件帮助类
func SendMail(user, password, host, to, subject, body, mailtype string) error {
	hp := strings.Split(host, ":")
	auth := smtp.PlainAuth("", user, password, hp[0])
	fmt.Println("1")
	var content_type string
	if mailtype == "html" {
		content_type = "Content-Type: text/" + mailtype + "; charset=UTF-8"
	} else {
		content_type = "Content-Type: text/plain" + "; charset=UTF-8"
	}
	msg := []byte("To: " + to + "\r\nFrom: " + user + "<" + user + ">\r\nSubject: " + subject + "\r\n" + content_type + "\r\n\r\n" + body)
	send_to := strings.Split(to, ";")
	fmt.Println(send_to)
	err := smtp.SendMail(host, auth, user, send_to, msg)
	fmt.Println("2")
	return err
}
//这里的端口有点不一样465/25都试一下
//err := SendMail("发送的邮箱", "发送的邮箱密码", "smtp.qq.com:25", "目标邮箱", "邮件标题", "邮件内容", "html")
func main() {
	err := SendMail("1558030788@qq.com", "zjvuwigwunurhfad", "smtp.qq.com:25", "z_chuanlin@163.com;1558030788@qq.com", "找回密码", "点击这里修改密码", "html")
	if err!=nil{
		log.Println(err)
	}
}
