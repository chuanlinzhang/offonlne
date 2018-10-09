package singleton

import (
	"testing"
	"fmt"
)

func TestGetIntance(t *testing.T) {
	l1:=GetIntance()
	l2:=GetIntance()
	fmt.Println(l1.name,l2.name)
}
