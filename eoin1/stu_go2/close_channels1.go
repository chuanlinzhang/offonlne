package main

import (
	"math/rand"
	"time"
	"log"
	"sync"
)

/*
M个接收者，一个发生者，发生者通过关闭数据通道说不在发送
 */
func main() {
	rand.Seed(time.Now().UnixNano())
	log.SetFlags(0)
	const MaxRandomNumber = 100000
	const NumReceivers   = 100
	wgReceivers:=sync.WaitGroup{}
	wgReceivers.Add(NumReceivers)

	dataCh:=make(chan int,100)
	//一个发送者
	go func() {
		for{
			if value:=rand.Intn(MaxRandomNumber);value==0{
				//一个发送者是关闭通道时安全的
				close(dataCh)
				return
			}else {
				dataCh<-value
			}
		}
	}()
	//多个接收者
	for i:=0;i<NumReceivers ;i++  {
		go func() {
			defer wgReceivers.Done()
			//接收值，直到通道被关闭
			//通道的值缓冲队列是空的
			for value:=range dataCh{
				log.Println(i,value)
			}
		}()
	}
	wgReceivers.Wait()
}