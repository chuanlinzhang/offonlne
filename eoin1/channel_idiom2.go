package main

import (
	"fmt"
	"time"
)
/*
它从指定通道中读取数据直到通道关闭，才继续执行下边的代码。很明显，另外一个协程必须写入 ch（不然代码就阻塞在 for 循环了），而且必须在写入完成后才关闭。
suck 函数可以这样写，且在协程中调用这个动作，
 */
func main() {
	suck1(pump1())
	time.Sleep(time.Second * 2)
}
func pump1() chan int {
	ch := make(chan int)
	go func() {
		for i := 0; ; i++ {
			ch <- i
		}
	}()
	return ch
}
func suck1(ch chan int) {
	go func() {
		for v := range ch {
			fmt.Println(v)
		}
	}()
}
