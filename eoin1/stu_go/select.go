package main

import (
	"fmt"
	"time"
)

func main() {
	//select语句属于条件分支流程控制语句，不过他只能用于通道。他可以包含若干条case语句。
	//并根据条件选择其中之一执行，select语句的case关键字只能后跟用于通道的发生操作的表达式或者
	//或者接受操作的表达式
	//golang的select的功能和select ，poll ，epoll相似。就是监听io操作，当io发生时
	//触发相应的动作
	var ch1=make(chan int)
	//生成一个协程
	go func() {
		for i:=0;i<3 ;i++  {
			ch1<-i
		}
	}()
	defer close(ch1)
	done:=0
	finshed:=0
	for finshed<3{
		select {
		case v,ok:=<-ch1:
			if ok{
				done=done+1
				fmt.Println(v)
			}
		}
		finshed = finshed+1
	}
	fmt.Println("Done",done)



	//当for和select结合使用时，break语言是无法跳出for之外的，因此如要break出来
	//这里需要加一个标签，使用goto，或者break到具体的位置
	//这里是使用break的样咧
	i:=0
	forend:
		for{
			select {
			case <-time.After(time.Second*2):
				i++
				if i==5{
					fmt.Println("break now")
					break forend
				}
				fmt.Println("inside the select")
			}
		}

		//使用goto
		i=0
		for{
			select {
			case <-time.After(time.Second*2):
				i++
				if i==5{
					fmt.Println("break now")
					goto ForEnd//需要在跳出结构的后面
				}
				fmt.Println("inside the select:")
			}
			fmt.Println("inside the for:")
		}
		ForEnd:
}

