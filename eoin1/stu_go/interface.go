package main

import "fmt"

//定义：interface 是一组抽象方法（未具体实现的方法/仅包含方法名，参数，返回值的方法）
// 的集合，有点像但又不同于其他编程语言的interface.type interfacennName interface{
// 方法列表}
//注意：1：interface可以被任意对象实现。一个类型/对象也可以实现多个interface
//2：方法不能重载，如 eat（） eat（s string）不能同时存在
//值：声明为interface类型的变量，可以储存任何实现了interface中所有方法的类型的变量（对象）
//类的值类型传递方法会自动生成对应的引用类型传递方法。反之不成立
//组合：将一个interface1嵌入到另一个interface2中声明中，其作用相当于把interface1中
//函数包含到interface2中，但是组合中不同，有重复的方法

type person struct {
	name string
	age  int
}

func (p person) printMsg() {
	fmt.Printf("i am %s and my age is %d.\n", p.name, p.age)
}

func (p person) eat(s string) {
	fmt.Printf("%s is eat %s\n", p.name, s)
}
func (p person) drink(s string) {
	fmt.Printf("%s is drinking %s\n", p.name, s)
}

type people interface {
	printMsg()
	eat(s string)
	drink(s string)
}

/*
任何类型都可以是interface
要点：1.interface 关键字用来定义一个接口，2.go 没有implements ，extend等关键字
3.实现一个接口的方法就是直接定义接口中方法，4.要实现多态纪要用指针和&object
 */

func main() {
	//定义一个people interface类型的变量p1
	var p1 people
	p1 = person{"eoin", 20}
	p1.printMsg()
	p1.drink("1111")
	//空接口类型 interface{} 使得我们可以向函数传递任意类型的变量
	//断言解决在使用interface{}的情况下，空接口类型像普通类型转换的类型转换问题
	//普通类型的之间的转换使用显示转换
	//空接口 ，无方法，及interface{}。可以代表任何类型的，有点类似于java中的object类

}
