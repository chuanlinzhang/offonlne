package simpleFactory

import "fmt"

/*
简单工厂模式:
简单工厂模式(Simple Factory Pattern)：又称为静态工厂方法(Static Factory Method)模式，
它属于类创建型模式（同属于创建型模式的还有工厂方法模式，抽象工厂模式，单例模式，建造者模式）。
在简单工厂模式中，可以根据参数的不同返回不同类的实例。
简单工厂模式专门定义一个类来负责创建其他类的实例，被创建的实例通常都具有共同的父类。
 */
/*
简单工厂模式
go 语言没有构造函数一说，所以一般会定义NewXXX函数来初始化相关类。 NewXXX 函数返回接口时就是简单工厂模式，也就是说Golang的一般推荐做法就是简单工厂。

在这个simplefactory包中只有API 接口和NewAPI函数为包外可见，封装了实现细节。
 */
//定义一个接口api
type API interface {
	Say(name string) string
}

func NewAPI(t int) API  {
	if t==1{
		return  &hiAPI{}
	}else if t==2{
		return  &helloAPI{}
	}
	return nil

}

//hiAPI 实现API接口
type hiAPI struct {

}
//实现接口方法
func (h *hiAPI) Say(name string) string  {
	return fmt.Sprintf("HI ,%s",name)
}
//helloAPI 实现API接口
type helloAPI struct {

}
//实现接口方法
func (h *helloAPI) Say(name string) string  {
	return fmt.Sprintf("Hello ,%s",name)
}
func main() {
	api:=NewAPI(1)//可以根据参数的不同返回不同类的实例
	say:=api.Say("eoin")
	fmt.Println(say)
}