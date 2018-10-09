package main

import (
	"fmt"
	"math/rand"
	"time"
)

func productor(ch chan<- string) {
	for {
		ch <- fmt.Sprint("%v", rand.Float64())
		time.Sleep(time.Second * 3)
	}
}
func customer(ch <-chan string) {
	for {
		message := <-ch
		fmt.Println(message)
	}
}
func main() {
	ch := make(chan string, 5)
	go productor(ch)
	customer(ch)

}
