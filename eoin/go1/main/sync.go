package main

import (
	"sync"
	"fmt"
	"time"
)
/*
其中Done()是Add(-1)的别名。简单的来说，使用Add()添加计数，Done()减掉一个计数，计数不为0, 阻塞Wait()的运行。
 */
func main() {
	var wg sync.WaitGroup
	for i:=0;i<10 ;i++  {
		for j:=0;j<10 ;j++  {
			wg.Add(1)
			go func() {
				fmt.Println("out",i,j)
				time.Sleep(time.Second*20)
				wg.Done()
			}()

		}
		wg.Wait()//当内层循环没有运行完之前，后再这里阻塞，等待内层运行完，在进入外层循环的下一次循环
	}
}
