package main

import (
	"os"
	"fmt"
)

/*
用切片读写文件,在控制台输出文件内容
 */
func main() {
	//打开需要的文件
	inputFile, inputError := os.Open("inpu.txt")
	if inputError != nil {
		fmt.Println("错误")
		return
	}
	//关闭文件
	defer inputFile.Close()
	cat(inputFile)//调用把文件写到控制台的函数，填入参数
}//文件写到控制台的函数
func cat(f *os.File) {
	const NBUF = 512
	var buf [NBUF]byte//作为流的缓冲区
	for {
		switch nr, err := f.Read(buf[:]); true {
		case nr < 0:
			fmt.Fprintf(os.Stderr, "%s", err.Error())
			os.Exit(1) //错误退出
		case nr == 0:
			return //退出函数
		case nr > 0:
			if nw, ew := os.Stdout.Write(buf[0:nr]); nw != nr {
				fmt.Fprintf(os.Stderr, "%s", ew.Error())
			}

		}
	}
}
