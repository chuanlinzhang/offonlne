package main

import "testing"

func TestAdd( t *testing.T)  {
	if 1>2{
		t.Log("ok")
	}
	if 1==2{
		t.Error()
	}
}
