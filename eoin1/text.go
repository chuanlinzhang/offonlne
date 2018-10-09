package main

import "fmt"

func main() {
	t()
}
func t() {
	if i := 1; i > 0 {

		fmt.Println("yes")
		//return //不会再执行后面的代码(包括defer)，跳出所在函数
		//defer fmt.Println("yes11")
		return
	}
	fmt.Println("no")
	defer fmt.Println("yes22")
}

/*
结论：defer确实是在return之前调用的。但表现形式上却可能不像。本质原因是return xxx语句并不是一条原子指令，
defer被插入到了赋值 与 RET之前，因此可能有机会改变最终的返回值
 */
//func f() (result int) {
//	defer func() {
//		result++
//	}()
//	return 0//值为1
//}
//func f() (result int) {
//	result = 0  //return语句不是一条原子调用，return xxx其实是赋值＋RET指令
//	func() { //defer被插入到return之前执行，也就是赋返回值和RET指令之间
//		result++
//	}()
//	return
//}
