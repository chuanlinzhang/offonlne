package main

import (
	"strings"
	"fmt"
)

func main() {
	st := strings.Split("03:13", ":")
	fmt.Println(st[0][1])
}
