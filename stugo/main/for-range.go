package main

import (
	"fmt"
	"time"
)

type a struct {
	b int
}

func main() {
	m:=make([]a ,0)
	m=append(m,a{b:1})
	m=append(m,a{b:1})
	m=append(m,a{b:1})
	for i:=range m {
		m[i].b=2
	}
	for _,x:=range m{
		fmt.Println(x)
	}
	rand01()

}

func rand01()   {
	ch:=make(chan int,10)
	go  func() {
		for {
			select {//select会尝试执行各个case, 如果都可以执行，那么随机选一个执行
			case ch<-0:
			case ch<-1:
			case <-time.NewTicker(time.Second*1).C:

			}
		}
	}()
	for k:=range ch{
		fmt.Println(k)
	}

}