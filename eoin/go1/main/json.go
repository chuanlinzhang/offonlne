package main

import (
	"encoding/json"
	"fmt"

	"os"
)



type People struct {
	Name     string
	Age      int
	Boy      bool
	Interest []string
}

var peoples = []People{
	{Name: "111", Age: 18, Boy: true, Interest: []string{"11", "22"}},
	{Name: "222", Age: 18, Boy: false, Interest: []string{"33", "33"}},
}

func main() {
	//data, err := json.Marshal(peoples)
	//if err != nil {
	//
	//}
	//fmt.Printf("%s\n", data)
	//data, err := json.MarshalIndent(peoples, "", "   ")

	//if err != nil {
		//log.Fatalf("JSON marshaling failed: %s", err)
	//}
	//fmt.Printf("%s\n", data)
	//file, _ := os.OpenFile("vcard.json", os.O_CREATE|os.O_WRONLY, 0666)

	//enc := json.NewEncoder(file) //文件流（写入到json文件）
	//err1 := enc.Encode(peoples)  //被实例化的结构体切片
	//if err != nil {
	//	panic(err1)
	//	return
	//}
	var po []People
	file,_:=os.Open("vcard.json")
	en := json.NewDecoder(file) //文件流（读取json文件）
	err2 := en.Decode(&po)       //只是定义，没有初始化或者实例化的结构体,记住这里要传地址
	defer file.Close()
	if err2 != nil {
		panic(err2)
		return
	}
	fmt.Println(po)

	//os.Stdin标准输入等价于把输入的内容封装成一个文件
}