package main

import "fmt"

func main() {
	var s []int

	s = append(s, 1, 2)
	fmt.Println(s)
	//s[2]=2//panic: runtime error: index out of range
	fmt.Println(s[0])
	fmt.Println(s)
}
//切片的基本用法和数组类似，但是含有下标操作，只能操作已有的下标，如果要增加元素使用append（）方法