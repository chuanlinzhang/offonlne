package main

import (
	"math/rand"
	"time"
	"log"
	"sync"
	"strconv"
)

/*
不要关闭接收端的频道，如果频道有多个并发发送者，请不要关闭频道。经典
m个接收方。n个发生方，随机的一个通过通知主持人不安比一个额外的信号通道说“
让我们结束游戏”
这是最复杂的情况，我们不能让任何接收者和发送者关闭数据通道，而且我们不能让任何
接收者关闭一个额外的信号通道来通知所以发送者和接收者退出游戏，
这两者都将打破通道关闭原则，但是，我们可以引入主持人的角色来关闭附加信号通道
 */
func main() {
	rand.Seed(time.Now().UnixNano())
	log.SetFlags(0)

	const MaxRandomNumber = 100000
	const NumReceivers = 10
	const NumSenders = 1000

	wgReceivers := sync.WaitGroup{}
	wgReceivers.Add(NumReceivers)

	dataCh := make(chan int, 100)
	stopCh := make(chan struct{})
	//stopCh他的发送者是下面主持人的协程里
	//他的接收者是数据的发送者和接收者
	toStop := make(chan string, 1)
	//toStop用于通知主持人
	//关闭附加的通道stopCh
	//他的发生者是任何发送者和接收者
	//他的接受者是主持人
	var stoppedBy string
	//主持人
	go func() {
		stoppedBy = <-toStop
		close(stopCh)
	}()
	//发送者
	for i := 0; i < NumReceivers; i++ {
		go func(id string) {
			for {
				value := rand.Intn(MaxRandomNumber)
				if value == 0 {
					//这里通知主持人，关闭附加通道
					select {
					case toStop <- "sender%" + id:
					default:
					}
					return
				}
				//
				select {
				case <-stopCh: //如果stopCh关闭，值为0
					return
				default:
				}
				//
				select {
				case <-stopCh: //如果stopCh关闭，值为0
					return
				case dataCh <- value:

				}

			}
		}(strconv.Itoa(i))
	}

	//接收者
	for i := 0; i < NumReceivers; i++ {
		go func(id string) {
			defer wgReceivers.Done()
			for { //这是第一个选择，要尽早的退出goroutine
				select {
				case <-stopCh:
					return
				default:
				}
				select {
				case <-stopCh:
					return
				case value := <-dataCh://接收数据
					if value == MaxRandomNumber-1 {
						select {
						case toStop <- "receiver#" + id:
						default:
						}
						return
					}
					log.Println(value)
				}
			}

		}(strconv.Itoa(i))
	}
	wgReceivers.Wait()
	log.Println("stopped",stoppedBy)
}
