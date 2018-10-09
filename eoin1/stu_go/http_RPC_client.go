package main

import (
	rpc2 "net/rpc"
	"log"
	"fmt"
)

type Params1 struct {
	Width,Height int
}

func main() {
	//连接远程rpc服务
	rpc,err:=rpc2.DialHTTP("tcp","127.0.0.1:8080")
	if err!=nil{
		log.Fatal(err)
	}
	ret:=0
	//调用远程方法，注意第三个参数是指针类型
	err=rpc.Call("Rect.Area",Params1{30,10},&ret)
	if err!=nil{
		log.Fatal(err)
	}
	fmt.Println(ret)

	err=rpc.Call("Rect.Perimeter",Params1{30,50},&ret)
	if err!=nil{
		log.Fatal(err)
	}
	fmt.Println(ret)
}
