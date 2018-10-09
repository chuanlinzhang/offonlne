package main

import (
	"fmt"
	"encoding/json"
	"os"
	"log"
)

type Address struct {
	Type    string
	City    string
	Country string
}

type VCard struct {
	FirstName string
	LastName  string
	Addresses []*Address
	Remark    string
}

func main() {
	pa := &Address{"peivate", "Aartselaar", "Belgium"}
	wa := &Address{"work", "Boom", "Belgium"}
	vc := VCard{"Jan", "Kersschot", []*Address{pa, wa}, "none"}
	fmt.Printf("%v:\n", vc)
	//{Jan Kersschot [0xc042076060 0xc042076090] none}:
	js, _ := json.Marshal(vc) //json格式化
	//{"FirstName":"Jan","LastName":"Kersschot","Addresses":[{"Type":"peivate","City":"Aartselaar","Country":"Belgium"},{"Type":"work","City":"Boom","Country":"Belgium"}],"Remark":"none"}
	fmt.Printf("%s", js)
	//是一个编译器,按照json的格式写到一个文件中
	file, _ := os.OpenFile("vcard.json", os.O_CREATE|os.O_WRONLY, 0666)
	defer file.Close()
	enc := json.NewEncoder(file)
	err := enc.Encode(vc)
	if err != nil {
		log.Println("Error in encoding json")
	}
	/*
	出于安全考虑，在 web 应用中最好使用 json.MarshalforHTML() 函数，其对数据执行HTML转码，所以文本可以被安全地嵌在 HTML <script> 标签中。

json.NewEncoder() 的函数签名是 func NewEncoder(w io.Writer) *Encoder，返回的Encoder类型的指针可调用方法 Encode(v interface{})，
	将数据对象 vc 的json编码写入 io.Writer file 中。
	 */
}
/*
json 包提供 Decoder 和 Encoder 类型来支持常用 JSON 数据流读写。NewDecoder 和 NewEncoder 函数分别封装了 io.Reader 和 io.Writer 接口。

func NewDecoder(r io.Reader) *Decoder
func NewEncoder(w io.Writer) *Encoder
要想把 JSON 直接写入文件，可以使用 json.NewEncoder 初始化文件（或者任何实现 io.Writer 的类型），
并调用 Encode()；反过来与其对应的是使用 json.Decoder 和 Decode() 函数：
 */