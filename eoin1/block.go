package main

import (
	"fmt"

	"time"
)

func f1(in chan int) {
	fmt.Println(<-in)
}

func main() {
	out := make(chan int)
	go f1(out)
	out <- 2//当向无缓冲通道放数据时，要确保接受者已经准备就绪。不然就会阻塞
time.Sleep(5*time.Second)

}
/*
缓冲容量和类型无关，所以可以（尽管可能导致危险）给一些通道设置不同的容量，只要他们拥有同样的元素类型。内置的 cap 函数可以返回缓冲区的容量。

如果容量大于 0，通道就是异步的了：缓冲满载（发送）或变空（接收）之前通信不会阻塞，元素会按照发送的顺序被接收。如果容量是0或者未设置，
通信仅在收发双方准备好的情况下才可以成功。
 */
 /*
 for i, v := range data {
	go func (i int, v float64) {
		doSomething(i, v)
		...
	} (i, v)
}
在 for 循环中并行计算迭代可能带来很好的性能提升。不过所有的迭代都必须是独立完成的
  */