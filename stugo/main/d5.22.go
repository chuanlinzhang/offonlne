package main

import (
	"fmt"
	"strconv"
	"strings"
	"encoding/binary"
	"math"
)

func main() {
var num1 float64=1.000
	var num2 float64=1.000
	var num3 float64=1111
	fmt.Println(powerFilter(num1,num2,num3))
//fmt.Println(num)
// str:=fmt.Sprintf("%v",num)
// fmt.Println(str)
// ss:=strings.Split(str,".")
// fmt.Println(len(ss[0]))
//	fmt.Println(len(ss[1]))
}
func powerFilter(c,m,t float64) int {
	var skwhslice,smcslice ,sthslice []string
	kwh:=c
	mc:=m
	th:=t
	skwh:=strconv.FormatFloat(kwh,'f',-1,64)
	smc:=strconv.FormatFloat(mc,'f',-1,64)
	sth:=strconv.FormatFloat(th,'f',-1,64)
	if strings.Contains(skwh,"."){
		skwhslice=strings.Split(skwh,".")
	}else {
		skwhslice=append(skwhslice,skwh)
		skwhslice=append(skwhslice,"")
	}
	if strings.Contains(smc,"."){
		smcslice=strings.Split(smc,".")
	}else {
		smcslice=append(smcslice,smc)
		smcslice=append(smcslice,"")
	}
	if strings.Contains(sth,"."){
		sthslice=strings.Split(sth,".")
	}else {
		sthslice=append(sthslice,sth)
		sthslice=append(sthslice,"")
	}
	skwhslice0:=len(skwhslice[0])
	skwhslice1:=len(skwhslice[1])
	smcslice0:=len(smcslice[0])
	smcslice1:=len(smcslice[1])
	sthslice0:=len(sthslice[0])
	sthslice1:=len(sthslice[1])
	fmt.Println(sthslice1)
	if kwh<0||skwhslice0>9||skwhslice1>2{
		return 2
	}
	if mc<0||smcslice0>9||smcslice1>2{
		return 3
	}
	if th<0||sthslice0>4||sthslice1>0{
		return 4
	}
	return 1
}
func xiao(num float64) int {
	smc:=strconv.FormatFloat(num,'f',0,64)
fmt.Println(smc)
	smc64,_:=strconv.ParseFloat(smc,64)
	fmt.Println(smc64)
	num=num-smc64
	fmt.Println(num)
	for i:=0;i<10 ; i++ {
		num=num*10
		if num-smc64==0{
			return i+1
		}
	}
	return 0
}
func Float64ToByte(float float64) []byte {
	bits := math.Float64bits(float)
	bytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(bytes, bits)

	return bytes
}