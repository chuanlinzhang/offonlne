package main

import (
	"bufio"
	"os"
	"fmt"
)

func main() {
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("请输入你的名字")
	input, err := inputReader.ReadString('\n')
	if err != nil {
		fmt.Println("错误")
		return
	}
	fmt.Println("你的名字是：", input)

	switch input{
	   case "s\n"://必须加上\n
		   fmt.Println("666")
	}
}
