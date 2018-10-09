package main

import (
	"net/http"
	"html/template"
	"fmt"
	"os"
	"io"
	"log"
)

//上传方法
func upload(w http.ResponseWriter,r *http.Request)  {
	//这里是get请求
	if r.Method=="GET"{
		t,_:=template.ParseFiles("upload.html")
		t.Execute(w,nil)
	}else if r.Method=="POST"{
		//服务端调用r.ParseMultiparForm ，把上传的文件存储在内存
		r.ParseMultipartForm(32<<20)
		//这里获取表单的uploadfile
		file,handler,err:=r.FormFile("uploadfile")
		if err!=nil{
			fmt.Println(err)
			return
		}
		defer file.Close()
		//这里打开文件buffer
		f,err:=os.OpenFile("./upload"+handler.Filename,os.O_WRONLY|os.O_CREATE,0666)//在当前目录
		if err!=nil{
			fmt.Println(err)
			return
		}
		defer f.Close()
		//把文件cp到指定目录下
		io.Copy(f,file)
		fmt.Fprintf(w,"%v",handler.Header)
		fmt.Fprintf(w,"上传成功")
	}else {
		fmt.Println("error")
	}
}
func main() {
	http.HandleFunc("/upload",upload)
	err:=http.ListenAndServe(":8080",nil)
	if err!=nil{
		log.Fatal("ListenAndServe:",err)
	}
}