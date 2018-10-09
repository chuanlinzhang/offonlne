package main

import (
	"os"
	"bufio"
	"io"
	"fmt"
)

func main() {
	inputFile, _ := os.Open("inpu.txt")//打开需要读取的文件
	outputFile, _ := os.OpenFile("input.txt", os.O_WRONLY|os.O_CREATE, 0666)//打开需要写入到的文件，如果没有则自动建立
	defer inputFile.Close()//文件打开后都需要关闭
	defer outputFile.Close()//文件打开后都需要关闭
	inputReader := bufio.NewReader(inputFile)//通过打开后的文件，得到读取缓冲器
	outputWriter := bufio.NewWriter(outputFile)//通过打开后的文件，得到写入缓冲器
	for {//
		inputString, _, readerError := inputReader.ReadLine()//一行一行读取
		if readerError == io.EOF { //io.EOF：文件已经读到结尾了,你还要它继续读而出现的错误
			fmt.Println("EOF")
			return
		}
		outputString := string(inputString[2:5]) + "\n"//将切片类型转化为字符串类型
		_, err := outputWriter.WriteString(outputString)//写入数据
		outputWriter.Flush()//在缓冲写入的最后千万不要忘了使用 Flush()，否则最后的输出不会被写入
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	fmt.Println("Conversion done")
}
