package main

import "fmt"

type people struct {
	name string
}

type user struct {
	name string
	eatt
}

type eatt interface {
	eat()
}

func (p people) eat()  {
	fmt.Println(p.name+"eat")
}
//func (u user) eat()  {
//	fmt.Println(u.name+"eat")
//}
//func tt( e eatt)  {
//	e.eat()
//}
func main() {
	p:=&people{name:"people"}
	u:=&user{eatt:p,name:"user"}//这里想调用已经实现这个接口的结构对应的方法，就可以用结构的实例来等价于这个接口
	//tt(p)
	u.eat()

}
//1.用接口来当作函数的参数传递，此时的形参可以传入所有实现这个接口的实例，更具不同的实例实现多太
//2.接口可以内嵌接口，相当于外部接口的拥有内部接口的所以方法、
//3.结构中内嵌接口，1.想调用已经实现这个接口的结构对应的方法，就可以用结构的实例来等价于这个接口
//2.另外本结构也可以重写接口的所以方法
//不能把接口拿来调用接口中的方法，因为接口是抽象，不能实例化，而调用方法必须要有一个载体，这个如果是函数也不行，实验过