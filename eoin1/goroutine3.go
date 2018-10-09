package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)

	go sendData(ch)
	go getData(ch)
	time.Sleep(time.Second)
}
func sendData(ch chan string) {
	ch <- "111111"
	ch <- "222222"
	ch <- "333333"
	ch <- "444444"
	ch <- "555555"
	ch <- "666666"
	ch <- "777777"
	//close(ch) //只有发送者才会关闭通道
}
func getData(ch chan string) {
	//for {
	//	//fmt.Printf("%s", <-ch)
	//	input, open := <-ch
	//	fmt.Println(open)
	//	if !open {
	//		break
	//	}
	//	fmt.Printf("%s", input)
	//}
	for input := range ch {
		fmt.Println(input)
	}

}
