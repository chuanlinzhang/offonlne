package main

import (
	"net/http"
	"fmt"
	"strings"
	"html/template"
	"log"
)

func sayHelloName(w http.ResponseWriter,r *http.Request)()  {
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("path",r.URL.Path)
	fmt.Println("scheme",r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k,v:=range r.Form{
		fmt.Println(k,strings.Join(v,""))
	}
	fmt.Fprintf(w,"%s","hello box")
}
func login(w http.ResponseWriter,r *http.Request)  {
	//打印请求的方法
	fmt.Println("method",r.Method)
	if r.Method=="GET"{//如果请求方法为get显示login.html，并相应给前端
	t,_:=template.ParseFiles("login.html")
	t.Execute(w,nil)
	}else {
		//否则就打印输出post接受的参数username和password
		fmt.Println(r.PostFormValue("username"))
		fmt.Println(r.PostFormValue("password"))
	}
}
func main() {
	//监听 走sayHelloName
	http.HandleFunc("/",sayHelloName)
	//路由控制/login 走login方法
	http.HandleFunc("/login",login)
	err:=http.ListenAndServe(":8080",nil)
	if err!=nil{
		log.Fatal("ListenAndServe:",err)
	}
}