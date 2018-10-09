package main

import (
	"time"
	"fmt"
)

func main() {
	the_time, _ := time.Parse("2006-01-02 15:04:05", "2015-01-01 11:38:56")
	fmt.Println(the_time)
	ctime:=the_time.AddDate(0,0,-1)
	fmt.Println(ctime)

}
