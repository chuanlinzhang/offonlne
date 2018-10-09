package main

import (
	"fmt"
	"os"
	"bufio"
)

func main() {
	//其不是写入一个文件，
	// 而是写入一个 io.Writer 接口类型的变量
	fmt.Fprintf(os.Stdout, "%s\n", "666")
	buf := bufio.NewWriter(os.Stdout)
	fmt.Fprintf(buf, "%s\n", "2222")
	buf.Flush()
}
