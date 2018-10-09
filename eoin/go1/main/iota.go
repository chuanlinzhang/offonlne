package main

import "fmt"

const (
	kb int64 = 1000
	mb       = kb * 1000
	gb       = mb * 1000
)
//Go的常量声明必须再编译时就确定值，类型。int64可以不要，
func main() {

	fmt.Println(kb, mb, gb)
}
