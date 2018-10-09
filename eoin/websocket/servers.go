package main

import (
	"fmt"
	"net/http"

	"golang.org/x/net/websocket"
	"log"
	"time"
)

func Echo(ws *websocket.Conn) {
	var err error

	for {
		var reply string
		if err = websocket.Message.Receive(ws, &reply); err != nil {
			fmt.Println("can't receive")
			break
		}

		fmt.Println("服务端接收到的消息" + reply)
		msg := "服务端发送" + reply
		fmt.Println(msg)
		 go func() {//每次来一个客户端的请求就开启一个协程来处理相应的请求
		 	for{
				time.Sleep(time.Second*4)
				if err = websocket.Message.Send(ws, msg); err != nil {
					fmt.Println("can't send")
					return
				}

			}

		 }()

	}
}
func main() {
	http.Handle("/", websocket.Handler(Echo))
	if err := http.ListenAndServe(":4001", nil); err != nil {
		log.Fatal("ListenAndServe", err)
	}
}
