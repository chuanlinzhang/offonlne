package main

import (
	"io/ioutil"
	"fmt"
)

func main() {

	inputFile := "inpu.txt"
	outoutFile := "inpucopy.txt"

	buf, err := ioutil.ReadFile(inputFile) //
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%s", string(buf))

	err = ioutil.WriteFile(outoutFile, buf, 322)
	if err != nil {
		panic(err.Error())
	}

}
