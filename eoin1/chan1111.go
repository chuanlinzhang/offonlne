package main

import (
	"fmt"
	"time"
)
/*
只接收的通道（<-chan T）无法关闭，
因为关闭通道是发送者用来表示不再给通道发送值了，所以对只接收通道是没有意义的。通道创建的时候都是双向的，
但也可以分配有方向的通道变量
 */
func main() {
	var c = make(chan int) // bidirectional
	go source(c)
	go sink(c)
  time.Sleep(2*time.Second)
}
func source(ch chan<- int) {
	for {
		ch <- 1
	}
}

func sink(ch <-chan int) {
	for {
		fmt.Println(<-ch)
	}
}
