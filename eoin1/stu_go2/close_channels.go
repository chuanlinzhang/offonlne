package main

import (
	"fmt"
)

/*
在golang中优雅的关闭频道
 */

type T int

func IsClosed(ch <-chan T) bool { //只读通道
	select {
	case <-ch://当通道关闭时也会触发这里
		fmt.Println(<-ch)//0，通道关闭
		return true
	default:
	}
	return false//通道没有关闭
}

//func SafeClose(ch chan T) (justClosed bool) {
//	defer func() {//如果处于任何原因从接受端或者通道的多个发生者之一关闭通道
//		if recover() != nil {//使用恢复机制来防止可能的恐慌，
//			justClosed = false
//		}
//	}()
//
//	close(ch)
//	return true
//}
func main() {
	c := make(chan T)

	fmt.Println(IsClosed(c))//false
	close(c)
	fmt.Println(IsClosed(c))//true
}
