package main

import (
	"net"
	"log"
	"fmt"
	"bufio"
)

func main() {
	listener,err:=net.Listen("tcp","localhost:8000")//监听这样一个请求
	if err!=nil{
		log.Fatal(err)
	}
go broadcaster()//用来接受通信息并进行相应的处理的函数
for{
	conn,err:=listener.Accept()//如果没有客户端连接请求，那么Accept函数会处于一个阻塞状态，直到有连接请求进来，创建一个连接
	if err!=nil{
		log.Fatal(err)
		continue
	}
	go handleConn1(conn)//为每一个连接，也就是每一个客户端开一个协程
}
}

type client chan<- string//只能发生数据的通道
var (
	entering = make(chan client)
	leaving = make(chan client)
	messages = make(chan string)//所以客户端的信息通道
)
/*
oadcaster监听来自全局的entering和leaving的channel来获知客户端的到来和离开事件。当其接收到其中的一个事件时，
会更新clients集合，当该事件是离开行为时，它会关闭客户端的消息发出channel。broadcaster也会监听全局的消息channel，
所有的客户端都会向这个channel中发送消息。
当broadcaster接收到什么消息时，就会将其广播至所有连接到服务端的客户端
 */
func broadcaster()  {
	clients:=make(map[client]bool)//所以客户端的连接,他的内部变量clients会记录当前建立连接的客户端集合。
	// 其记录的内容是每一个客户端的消息发出channel的"资格"信息
	for{
		select {
		case msg:=<-messages:
			//向所以客户端广播消息
			//遍历客户端，并通过每个客户端自己的通道给我自己发生消息
			for cli:=range clients{
				cli <-msg//注意，这里的cli等价于ch，cli是ch的一个引用
			}
		case cli:=<-entering://初始化每次连进来的客户端，在clients中存入
			clients[cli]=true
		case cli:=<-leaving://。客户端的读取方循环会在broadcaster接收到leaving通知并关闭了channel后终止。
			delete(clients,cli)
			close(cli)
		}
	}
}
func handleConn1(conn net.Conn)  {

	ch:=make(chan string)
	go clientWriter(conn,ch)
//到来的客户端
	who:=conn.RemoteAddr().String()//获取请求路径中路径片段
	ch <-"you are "+who
	messages<-who+"has arrived"
	entering<-ch
//发生的信息
	input:=bufio.NewScanner(conn)//从控制太输入的信息.输入流
	fmt.Println("ok"+"****")
       //fmt.Println(input.Scan())

	for input.Scan(){
		fmt.Println("ok")
		messages<-who+":"+input.Text()//输入信息
	}
//离开的客户端
	leaving<-ch
	messages<-who+":"+"has left"
	conn.Close()
}
/*
接收向客户端发出消息channel中发送的广播消息，并将它们写入到客户端的网络连接
 */
func clientWriter(conn net.Conn,ch <-chan string){
    for msg:=range ch{
    	fmt.Fprintln(conn,msg)
	}
}