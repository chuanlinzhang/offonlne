package main



/*
原型模式
原型模式使对象能复制自身，并且暴露到接口中，使客户端面向接口编程时，不知道接口实际对象的情况下生成新的对象。

原型模式配合原型管理器使用，使得客户端在不知道具体类的情况下，通过接口管理器得到新的实例，并且包含部分预设定配置。
 */
import (
	"fmt"
	"testing"
)

type Type1 struct {
	name string
}

func (t *Type1) Clone() Cloneable {

	return t
}
func TestPrototypeManager_Set(t *testing.T) {
	tt := Type1{
		name: "eoin",
	}
	pro := NewPrototypeManager()
	pro.Set(tt.name, &tt)
	ttt := pro.Get(tt.name)
	fmt.Println(ttt.(*Type1).name)
	fmt.Println(tt)
}
//注意：当结构体实现一个接口时，接受者用的值或者指针，后面的所有与接口相关处，出现的结构体，（值或者指针），类型保持一致的
