package users

import (
	"testing"
	"fmt"
)
//自动创建test数据库
func TestRegister1(t *testing.T) {
	customers := &Customers{}
	customers.LoginName="1"
	if b:=Register1(*customers);b{
		t.Log("ok")
	}else {
		t.Log("no")
		fmt.Println(*customers)
	}
}
func BenchmarkRegister1(b *testing.B) {
	customers := &Customers{}
	customers.LoginName = "1"
	Register1(*customers)

}