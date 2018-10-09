package main

import (
	"os"
	"fmt"
	"io"
	"io/ioutil"
	"bufio"
)

//用户错误检查的函数
func check(e error) {
	if e != nil {
		panic(e)
	}
}

/*
判断文件是否存在，存在返回true，不存在返回false
 */
func checkFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) { //os.IsNotExist(err)判断错误是否存在
		exist = false
	}
	return exist
}
func main() {
	var wireteString = "测试n"
	var filename = "./output1.txt"
	var f *os.File
	var err1 error
	var input string
	//第一种方式，使用io.writeString 写入文件
	if checkFileIsExist(filename) { //如果文件存在
		f, err1 = os.OpenFile(filename, os.O_APPEND, 0777) //打开文件，os.O_APPEND在写如文件时把数据追加到后面
		fmt.Println("文件存在")
	} else {
		f, err1 = os.Create(filename) //文件不存在，创建文件
		fmt.Println("文件不存在")
	}
	check(err1)
	n, err1 := io.WriteString(f, wireteString)//这种写入不会覆盖原来的数据（字符串）,后面三种会覆盖原来的数据
	check(err1)
	fmt.Printf("%d 字节", n)
	//第二种方式。使用ioutill.WriteFile写入文件（）
	var d1 = []byte(wireteString)//把字符串转换为字节
	err2:=ioutil.WriteFile("./output2.txt",d1,0666)//写入文件（字节数组）
	check(err2)
	//第三种 使用Fileal（Writer，WriteString）写文件
	f,err3:=os.Create("./output3.txt")//创建文件
	check(err3)
	defer  f.Close()
	n2,err3:=f.Write(d1)
	check(err3)
	fmt.Println(n2)
	f.Sync()
	//第四中
	w:=bufio.NewWriter(f)//var f *os.File
	w.WriteString("64511")
	w.Flush()
	f.Close()
	fmt.Scanf("%s",&input)//控制台输入
	fmt.Println(input)
}
