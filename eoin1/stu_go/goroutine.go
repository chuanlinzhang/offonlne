package main

import (
	"fmt"
	"time"
	"math/rand"
)

func p() {
	for i := 0; i < 2; i++ {
		fmt.Println(i)
		time.Sleep(time.Second * 1)
	}
}

func sell(c chan int) {
	for {
		num := <-c
		fmt.Println("sell", num, "bread")
	}
}

func produce(c chan int) {
	for {
		num := rand.Intn(10)
		t := time.Duration(num)
		fmt.Println("product", num)
		c <- num
		time.Sleep(t * time.Second)
	}
}
func main() {
	go p()
	var input string
	fmt.Scanln(&input)
	fmt.Println("end")
	//goroutine使用channel来通信
	c := make(chan int)
	go sell(c)
	go produce(c)
	var input2 string
	fmt.Scanln(&input2)
	fmt.Println("end")
}
