package main

import (
	"net/http"
	"fmt"
	"io/ioutil"

)

func main() {
	resp,err:=http.Get("http://www.baidu.com")
	if err!=nil{
		fmt.Println("get错误")
	}
	defer resp.Body.Close()
	body,_:=ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}
