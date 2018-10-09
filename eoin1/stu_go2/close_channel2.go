package main

import (
	"math/rand"
	"time"
	"log"
	"sync"
)

/*
一个接收器，n个发生器，接收器通过关闭另一个信号通道说“请停止发生更多”

我们不能让接收者关闭数据通道，因为这样会打破通道关闭原则，但是我们可以让接收器
关闭一个人额外的信号通道来通知发生器停止的发送值
 */

func main() {
	rand.Seed(time.Now().UnixNano())
	log.SetFlags(0)

	const MaxRandomNumber = 100000
	const NumSenders = 1000
	wgReceives := sync.WaitGroup{}
	wgReceives.Add(1)

	dataCh := make(chan int, 100)
	stopCh := make(chan struct{}) //额外加的单一通道
	//stopCh的发送方dataCh的接收方，反之

	//发送者
	for i := 0; i < NumSenders; i++ {
		go func() {
			//这里的第一个选择是尝试推出协程
			//可以省略的
			for {
				select {
				case <-stopCh:
					return
				default:

				}

				//
				select {
				case <-stopCh://当stopCh被关闭时，<-stopCh值为0，会执行这个case分支
					return
				case dataCh<-rand.Intn(MaxRandomNumber):

				}
			}
		}()
	}
	//一个接受者
	go func() {
		defer wgReceives.Done()
		for value:=range dataCh{
			if value==MaxRandomNumber-1{
				//这里接收了通道dataCh，也发生给stopCh
				//这里关闭也发生给stopCh通道时安全的
				close(stopCh)
				return
			}
			log.Println(value)
		}
	}()
	wgReceives.Wait()
}
//对于附加的通道，其发生者是数据通道的接受者，附加信号通道有唯一的发送者关闭它
//该发送者保持通道关闭原则
//这个列子中，通道dataCh永远不会关闭，是的，通道不必关闭，如果没有goroutine引用它
//那么通道最终将被垃圾收集，无论它是否关闭