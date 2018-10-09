package goconvey

import (
	"testing"
	"github.com/smartystreets/goconvey/convey"
)

func TestStringSliceEqual(t *testing.T) {
	convey.Convey("TestStringSliceEqual should return true when a != nil  && b != nil", t, func() {
	a := []string{"hello", "goconvey"}
	b := []string{"hello", "goconvey"}
	StringSliceEqual(a, b)

})
}
