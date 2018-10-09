package main

import (
	"net/http"
	"fmt"
	"html/template"
	_"strconv"
	"log"
	"strconv"
)

func sayHelloName1( w http.ResponseWriter,r *http.Request)  {
	fmt.Fprintf(w,"%s","hello box")
}
func login1(w http.ResponseWriter,r *http.Request)  {
	if r.Method=="GET"{
		t,_:=template.ParseFiles("login1.html")
		t.Execute(w,nil)
	}else if r.Method=="POST"{//必须是大写
		username:=r.FormValue("username")
		password:=r.FormValue("password")
		like:=r.FormValue("like")
		sex:=r.FormValue("sex")
		utype:=r.FormValue("utype")

		fmt.Println(username,password,like,sex,utype)
		//获得年龄之后转成int型
		age,err:=strconv.Atoi(r.FormValue("age"))
		if err!=nil{
			w.Write([]byte("数字化出错"))
			return
		}
		if username==""||password==""||age==0{
			w.Write([]byte("not null"))
			return
		}
		//获取数据判定大小
		if age>100 {
			w.Write([]byte("age is to big"))
			return
		}

	}else {
 fmt.Println("error")
	}
}
func main() {
	http.HandleFunc("/",sayHelloName1)
	http.HandleFunc("/login",login1)
	err:=http.ListenAndServe(":8080",nil)
	if err!=nil{
		log.Fatal("ListenAndServe:",err)
	}
}