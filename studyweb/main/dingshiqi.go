package main

import (
	"time"
	//"log"
	"fmt"
)

func main() {
	//for {
	//
	//	t := time.NewTimer(time.Second * 10)
	//	log.Printf("时间为[%v]\n", time.Now())
	//
	//	expire := <-t.C
	//	fmt.Println("Expiration time: %v.\n", expire)
	//
	//}

	/*
	time.After函数
time.After函数， 表示多少时间之后，但是在取出channel内容之前不阻塞，后续程序可以继续执行

鉴于After特性，其通常用来处理程序超时问题
	 */
	ch1 := make(chan int, 1)
	ch2 := make(chan int, 1)
	select {
	case e1 := <-ch1:
		fmt.Println(e1)
	case e2 := <-ch2:
		fmt.Println(e2)
	case <-time.After(5 * time.Second):
		fmt.Println("out")


	}
}
