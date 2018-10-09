package main

import "fmt"

func main() {
	s := []int{1, 2, 3}
	ss := s[1:]
	for i, v := range ss {
		fmt.Println(i, v)
		ss[i] += 10//等现在切片的下标索引还在开始的底层数值范围内，
		           // 那么现在对切片的修改就是
		             // 会修改底层数组，
	}
	fmt.Println(s)
	fmt.Println(ss)
	ss = append(ss, 4)
	for i, v := range ss {
		fmt.Println(i, v)
		ss[i] += 10
	}
	fmt.Println(s)
	fmt.Println(ss)
}
