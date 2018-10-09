package main

import (
	"net/rpc"
	"net/http"
	"log"
)

//go对rpc的支持，支持三个级别：tcp，http，jsonrpc
//go的rpc只支持go开发的服务器与客户端的交互，因为采用了gob编码

type Params struct {
		Width,Height int
}
type Rect struct {
}
//求面积、
func (r *Rect) Area(p Params,ret *int)  error{
	*ret = p.Height*p.Width
	return nil
}
//求长度
func (r *Rect) Perimeter( p Params,ret *int) error  {
	*ret =(p.Width+p.Height) *2
	return nil
}

func main() {
	ret:=new(Rect)//
	//注册一个人rect服务
	rpc.Register(ret)
	//把服务绑定在http协议上
	rpc.HandleHTTP()
	err:=http.ListenAndServe(":8080",nil)
	if err!=nil{
		log.Fatal(err)
	}
}