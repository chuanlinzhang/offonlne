package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	go pump(ch1)
	go suck(ch1)
	time.Sleep(time.Second)
	//fmt.Println(<-ch1)//输出只有0
}
//给程序 1 秒的时间来运行：输出了上万个整数。
//通道提供数值。叫做生产者
func pump(ch chan int) {
	for i := 0; ; i++ {
		ch <- i
	}
}
func suck(ch chan int) {
	for {
		fmt.Println(<-ch)
	}
}
