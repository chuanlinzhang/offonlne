package main

import (
	"os"
	"fmt"
	"bufio"
	"io"


)

func main() {
	//打开文件
	inputFile, inputError := os.Open("inpu.txt")
	if inputError != nil {
		fmt.Println("错误")
		return
	}
	defer inputFile.Close()

	outputFile, err := os.Create("1.txt")//创建一个写入文件
	if err!=nil{
		fmt.Println(err)
	}
	defer outputFile.Close()
	wBuf := bufio.NewWriter(outputFile)//得到一个写入器变量
	//获得一个读取器变量
	inputReader := bufio.NewReader(inputFile)
	for {
		inputString, readerError := inputReader.ReadString('\n')
		wBuf.WriteString(inputString)//写文件
		fmt.Printf("%s", inputString)
		if readerError == io.EOF { //判断是否读到最后一行，然后退出循环退出循环
			fmt.Println(readerError) //EOF
			fmt.Println(io.EOF)      //EOF
			wBuf.Flush()
			return
		}


	}
}
