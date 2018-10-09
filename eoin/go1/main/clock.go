package main

import (
	"net"
	"log"
	"io"
	"time"
)

//一个顺序执行的时钟服务器，它会每隔一秒钟将当前的时间写到客户端
func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)//我们这里对服务端程序做一点小改动，使其支持并发：在handleConn函数调用的地方增加go关键字，
		// 让每一次handleConn的调用都进入一个独立的goroutine。可以有很多个客户端从这个服务端接受信息
	}

}
func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return
		}
		time.Sleep(time.Second * 1)
	}

}
/*
Listen函数创建了一个net.Listener的对象，这个对象会监听一个网络端口上到来的连接，
在这个例子里我们用的是TCP的localhost:8000端口。listener对象的Accept方法会直接阻塞，
直到一个新的连接被创建，然后会返回一个net.Conn对象来表示这个连接。

handleConn函数会处理一个完整的客户端连接。在一个for死循环中，用time.Now()获取当前时刻，
然后写到客户端。由于net.Conn实现了io.Writer接口，我们可以直接向其写入内容。这个死循环会一直执行，
直到写入失败。最可能的原因是客户端主动断开连接。这种情况下handleConn函数会用defer调用关闭服务器侧的连接，
然后返回到主函数，继续等待下一个连接请求。
 */