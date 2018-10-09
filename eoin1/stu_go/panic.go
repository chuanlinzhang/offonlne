package main

import "fmt"

func main() {
	//defer所在函数执行完所以代码之后，会自动执行defer函数，但是defer在
	// return之前执行，就是说defer函数可能会影响返回值
	//给一个列子，获取数组元素，处理数组访问越界的问题
	a:=[5]int{1,2,3,4,5}
	for i:=0;i<10 ;i++  {
		itme,ok:=get(i,a)
		fmt.Println("a[%d]=%d[%v]\n",i,itme,ok)
	}
	//recover 相当于try-catch部分，使得panic不再传递，而defer相当于try-catch
	//-final的fianl部分
	defer func() {
		fmt.Println("a")
		if err:=recover();err==nil{
			fmt.Println(err)
		}
	}()
	fmt.Println("b")
	panic("test")
	fmt.Println("c")
}
func get(i int,a [5]int)  (ret int,ok bool) {
	ok=true
	defer func() {
		err:=recover()
		if err!=nil{
			fmt.Println(err,"-1")
			ret=-1
			ok=false
		}
	}()
	ret=a[i]
	return
}