package main

import "fmt"

/*
工厂方法模式
工厂方法模式使用子类的方式延迟生成对象到子类中实现。

Go中不存在继承 所以使用匿名组合来实现
 */
//Operator 是封装的实际类接口
type Operator interface {
	SetA( int)
	SetB( int)
	Result() int
}
//OperatorFactory 是工厂接口
type OperatorFactory interface {
	Create() Operator
}
//OperatorBase 是Operator接口实现的基类，封装公用方法，还有一层加上私有的方法才产品实现类
type OperatorBase struct {
	a,b int
}
func (o *OperatorBase) SetA(a int)  {
	o.a=a
}
func (o *OperatorBase) SetB(b int)  {
	o.b=b
}
//PlusOperatorFactory 是 PlusOperator 的工厂类
type PulsOperatorFactory struct {
}
type PulsOperator struct {
	*OperatorBase
}
func (p *PulsOperator) Result() int {
	return p.a+p.b
}
func (PulsOperatorFactory) Create() Operator {
	return &PulsOperator{
		OperatorBase:&OperatorBase{},
	}
}
//MinusOperatorFactory 是 MinusOperator 的工厂类
type MinusOperatorFactory struct{}
//MinusOperator Operator 的实际减法实现
type MinusOperator struct {
	*OperatorBase
}
//Result 获取结果
func (o MinusOperator) Result() int {
	return o.a - o.b
}
func (MinusOperatorFactory) Create() Operator {
	return &MinusOperator{
		OperatorBase: &OperatorBase{},
	}
}
//计算
func compute(factory OperatorFactory,a,b int)  int {
	op:=factory.Create()
	op.SetA(a)
	op.SetB(b)
	return op.Result()
}
func main() {
	PulsOperatorFactory:=&PulsOperatorFactory{}//具体的工厂类
	re:=compute(PulsOperatorFactory,1,2)
	fmt.Println(re)
	MinusOperatorFactory:=&MinusOperatorFactory{}
	re=compute(MinusOperatorFactory,1,2)
	fmt.Println(re)
}
/*
工厂方法模式：
工厂接口。工厂接口是工厂方法模式的核心，与调用者直接交互用来提供产品。在实际编程中，有时候也会使用一个抽象类来作为与调用者交互的接口，其本质上是一样的。
工厂实现。在编程中，工厂实现决定如何实例化产品，是实现扩展的途径，需要有多少种产品，就需要有多少个具体的工厂实现。
产品接口。产品接口的主要目的是定义产品的规范，所有的产品实现都必须遵循产品接口定义的规范。产品接口是调用者最为关心的，产品接口定义的优劣直接决定了调用者代码的稳定性。同样，产品接口也可以用抽象类来代替，但要注意最好不要违反里氏替换原则。
产品实现。实现产品接口的具体类，决定了产品在客户端中的具体行为。
 */