package main

import (
	"sync"
	"fmt"
)

func main() {
	var wg sync.WaitGroup
	   i:=0
	for ;i<10 ;i++  {
		wg.Add(1)
		go func(i int ) {
			fmt.Println("hello ",i)
			wg.Done()
		}(i)
	}
	wg.Wait()//等待所有的协程结束，在执行后面的语句
	fmt.Printf("run",i)
}
