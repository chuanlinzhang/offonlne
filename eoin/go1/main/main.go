package main

import "fmt"

func main() {

	 pp :=po{"111"}
	fmt.Println(B(&pp))
}

type Aa interface {
	Add() string
}

func B(aa Aa) string {
	return  aa.Add()
}

type po struct {
	name string
}

func (p *po) Add()  string {
	return  p.name
}
