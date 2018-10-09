package main

import (
	"net"
	"strings"
	"github.com/astaxie/beego/logs"
	"time"
	"fmt"
)

func main() {
	server,err:=net.Listen("tcp",":1208")
	if err!=nil{
		logs.Error(err)
	}
	logs.Info("server started")
	for {
		conn,err:=server.Accept()
		if err!=nil{
			logs.Error(err)
			break
		}
		go connHandler(conn)
	}
}
//心跳计时,长链接超时设置
func HeartBeating(c net.Conn,cnt chan int)  {
	select {
	case <-cnt:
		c.SetReadDeadline(time.Now().Add(time.Duration(10)*time.Second))
	case <-time.After(time.Duration(15)*time.Second):
		logs.Info("nothing")
	}
}
func connHandler(c net.Conn)  {
	if c==nil{
		logs.Info("conn is null")
		return
	}
	buf:=make([]byte,4096)
	for {
		cnt,err:=c.Read(buf)
		logs.Error(err)
		fmt.Println(cnt)
		readslice:=make(chan int,1)
		readslice <-cnt
		go HeartBeating(c,readslice)
       //fmt.Println(<-readslice)

		if err!=nil||cnt==0{
			c.Close()
			logs.Error(err)
			break
		}
		instr:=strings.TrimSpace(string(buf[0:cnt]))
		inputs:=strings.Split(instr," ")
		fmt.Println(inputs)
		switch inputs[0] {
		case "ping":
			c.Write([]byte("pong\n"))
		case "echo":
			echostr:=strings.Join(inputs[1:]," ")+"\n"
			c.Write([]byte(echostr))
		case "quit":
			c.Close()
			break
		default:
			logs.Info("unsupported command")
			c.Write([]byte("unsupported command"))
		}
	}
	logs.Info("connction form close")
}