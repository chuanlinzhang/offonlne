package main

import "fmt"

func main() {
	var s []int
	for i:=1;i<=3;i++{
		s=append(s,i)
		reverse(s)
		fmt.Println(s)
	}

}
func reverse(s []int)  {
	for i,j:=0,len(s)-1;i<j ;i++  {
		j=len(s)-(i+1)
		s[i],s[j]=s[j],s[i]
	}
}
/*
slice是由一个指向数组指针支持的，这意味着即使我们的slice正在按值传递
数组指针依然指向相同的内存地址
 */