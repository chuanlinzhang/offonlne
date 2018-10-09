package main

import (
	"os"
	"strings"
	"fmt"
)

func main() {
	who := "4444"
	if len(os.Args) > 1 {
		who += strings.Join(os.Args[1:], " ")
	}
	fmt.Println("666", who)
}
