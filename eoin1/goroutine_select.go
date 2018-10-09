package main

import (
	"fmt"
	"time"
)
/*
有 2 个通道 ch1 和 ch2，三个协程 pump1()、pump2() 和 suck()。这是一个典型的生产者消费者模式。
在无限循环中，ch1 和 ch2 通过 pump1() 和 pump2() 填充整数；suck() 也是在无限循环中轮询输入的，
通过 select 语句获取 ch1 和 ch2 的整数并输出。选择哪一个 case 取决于哪一个通道收到了信息。
程序在 main 执行 1 秒后结束。
 */
func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go pump11(ch1)
	go pump22(ch2)
	go suck11(ch1, ch2)
	time.Sleep(time.Second)
}
func pump11(ch chan int) {
	for i := 0; ; i++ {
		ch <- i
	}
}
func pump22(ch chan int) {
	for i := 0; ; i++ {
		ch <- i
	}
}
func suck11(ch1, ch2 chan int) {
	for {
		select {
		case v := <-ch1:
			fmt.Printf("Received:%d\n", v)

		case v := <-ch2:
			fmt.Printf("Received:%d\n", v)
		}
	}
}
