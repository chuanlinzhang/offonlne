package main

import "fmt"

func main() {
	fmt.Println("1")
	for {
		Test()
	}

	fmt.Println("4")
}

func Test()  {
	//如果一个函数的中可能会出现panic，就用recover来捕获，记录错误，后跳出本函数，但本函数外的语句会继续执行，不会导致程序崩溃
	//如果没有recover，一旦某个地方出现panic，整个服务就会挂掉
	defer func() {
		fmt.Println("a")
		if err :=recover();err!=nil{
			fmt.Println(err)
		}
		fmt.Println("b")
	}()
	fmt.Println("2")
	panic("bug")
	fmt.Println("3")
}