package main

import (
	"net"
	"log"
	"io"
	"os"
)

func main() {
	conn,err:=net.Dial("tcp","localhost:8000")//程序会从连接中读取数据，并将读到的内容写到标准输出中
	if err!=nil{
		log.Fatal(err)
	}
	defer conn.Close()
	mustCopy(os.Stdout,conn)
}
func mustCopy(dst io.Writer,src io.Reader){
	if _,err:=io.Copy(dst,src);err!=nil{
		log.Fatal(err)
	}
}