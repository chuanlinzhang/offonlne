package main

import (
	"net/http"
	"fmt"
	"strings"
	"log"
)

func sayHelloName2(w http.ResponseWriter,r *http.Request)  {
	r.ParseForm()//解析form
	fmt.Println(r.Form)
	fmt.Println("path:",r.URL.Path)
	fmt.Println("scheme",r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k,v:= range r.Form{
		fmt.Println("key:",k)
		fmt.Println("val:",strings.Join(v,""))
	}
	fmt.Fprintf(w,"hello xx")
}
func main() {
	//路由
	http.HandleFunc("/",sayHelloName2)
	err:=http.ListenAndServe(":8080",nil)
	if err!=nil{
		log.Fatal(err)
	}
}