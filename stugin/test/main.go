package main

import (
	"os"
	"fmt"
)

func main() {
	test()
}
func test()  {
	file, _:= os.Open("studygolang.txt")
	fmt.Println(file)
	defer file.Close()
	//if err!=nil{
	//	fmt.Println(err)
	//	//return
	//}
	//file.Close() 会 panic，因为 file 是 nil。
	// 因此，应该将 defer file.Close() 放在错误检查之后。
	fmt.Println("666")
}