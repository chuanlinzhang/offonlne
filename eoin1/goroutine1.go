package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("main()")
	go longWait()
	go shortWait()
	fmt.Println("to sleep in main()")
	time.Sleep(10 * time.Second)
	fmt.Println("end main()")
}
func longWait() {
	fmt.Println("Beginning longWait()")
	time.Sleep(5 * time.Second)
	fmt.Println("end of longWait()")
}
func shortWait() {
	fmt.Println("Beginning shortWait()")
	time.Sleep(2 * time.Second)
	fmt.Println("End of shortwWait")
}
