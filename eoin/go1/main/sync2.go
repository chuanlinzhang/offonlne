package main

import (
	"time"
	"fmt"
	"sync"
)

/*
用于测试 给chan发送 1千万次，并接受1千万次的性能。
 */

const num  =100000000

func main() {
	TestFunc("test",TestChan)
}
func TestFunc(name string,f func())  {
	st:=time.Now().UnixNano()
	f()
	fmt.Printf("task %s cost %d \r\n",name,(time.Now().UnixNano()-st)/int64(time.Millisecond))
	
}
func TestChan()  {
	var wg sync.WaitGroup
	c:=make(chan string)
	wg.Add(1)
	go func() {
		for j:=range c{
			fmt.Println(j)
		}
		wg.Done()
	}()
	for i:=0;i<num ; i++ {
		c<-"123"
	}
	close(c)
	wg.Wait()
}