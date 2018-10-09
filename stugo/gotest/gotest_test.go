package gotest

import "testing"
func BenchmarkDivision(b *testing.B) {
	for i:=0;i<b.N ;i++  {
		Division(4,5)
	}
}
func BenchmarkDivision2(b *testing.B) {
	b.StopTimer()
	b.StartTimer()
	for i:=0;i<b.N ;i++  {
		Division(4,5)
	}
}
