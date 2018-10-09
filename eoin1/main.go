package main

import (
	"bufio"
	"os"
	"fmt"
)
//读取用户的键盘（控制台）输入
func main() {
	inputReader := bufio.NewReader(os.Stdin)
	input, err := inputReader.ReadString('\n')
	fmt.Println(input, err)
	//现在开始使用这种写法
}
