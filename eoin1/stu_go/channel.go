package main

import (
	"time"
	"fmt"

)

var c chan int  //c为nil 永久阻塞
/*var ccc int//基本数据类型声明系统会初始化，引用数据类型，需要使用make来开辟内存，初始化*/
var cc = make(chan int, 0)

func ready(w string, sec int) {
	time.Sleep(time.Duration(sec) * time.Second) //类型相同的才可以进行计算，这里将普通的int类型转换为时间对应的int类型time.Duration
	fmt.Println(w, "is ready")
	//往channel中插入数据
	cc <- 1
}

//有缓存的channel：这里有个缓存，因此放入数据的操作从<-0 先于取数据的操作<-c
var a string
var c1 = make(chan int, 10)

func f() {
	a = "hello world"
	c1 <- 0
}

//无缓冲的channel ：由于从c是无缓冲的channel ，因此必须保证取操作<-c 先于放操作c<-0
var a2 string
var c2=make(chan int)

func f2()  {
	a2="hellow my world"
	<-c2
}
func main() {
	fmt.Println(c)
	fmt.Println(cc)
	go func() {
		fmt.Println(<-cc)
	}()
	ready("111", 1)
   time.Sleep(time.Second*2)

  /*
  go语言提供的消息通信机制被称为channel ”不要通过共享内存来通信，而应该通过通信来共享内存“
  channel是go语言在语言级别提供的goroutine间的通信方式，channel是类型相关的chan 声明
  var chanName chan ElementType 并使用直接的内置函数make（）创建

   */
/*
总结：
  channel分为两种：一种是有buffer ，一种是没有buffer的 ，默认是没有buffer
cj：=make（chan int）
cj：=make（chan int ，0）
cs：=make（chan int ，0）
有缓冲的channel，因此要注意“放”先于“取“
无缓冲的channel，因此要注意”取“先于”方法“

关于chnnel的4个特性、
1.给一个nil channel 发送数据，造成永久阻塞
2.从一个nil channel 接受数据，造成永久阻塞
3.给一个已经关闭的channel发生数据，引起panic
4.从一个已经关闭的channel接收数据，立即返回0值

 */

}
