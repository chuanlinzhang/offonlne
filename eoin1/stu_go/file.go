package main

import (
	"os"
	"fmt"

)

func main() {
	//创建一个空文件
	myFile :="mytest.txt"
	file,err:=os.Create(myFile)
	if err!=nil{
		fmt.Println(err)
		os.Exit(0)
	}
	for i:=0;i<5 ;i++  {
		file.WriteString("this is my file\r\n")
		file.Write([]byte("file end\r\n"))
	}
	//这里打开上面创建的文件
	fin,err:=os.Open(myFile)
	if err!=nil{
		fmt.Println(err)
		os.Exit(0)
	}
	buf:=make([]byte,1024)
	for{
		//读取文件每次的字节数
		n,_:=fin.Read(buf)

		if 0==n{
			break
		}
		//标准输出所读文件
		os.Stdout.Write(buf[:n])
	}
	defer func() {
		file.Close()
		fin.Close()
		//删除文件
		del:=os.Remove(myFile)
		if del!=nil{
			fmt.Println(del)
			os.Exit(0)
		}
	}()
}
