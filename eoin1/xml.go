package main

import (
	"encoding/xml"
	"strings"
	"fmt"
)

var t, token xml.Token
var err error

func main() {
	input := "<Person><FirstName>Laura</FirstName><LastName>Lynn</LastName></Person>"
	inputReader := strings.NewReader(input)
	p := xml.NewDecoder(inputReader)
	for t, err = p.Token(); err == nil; t, err = p.Token() {
		switch token := t.(type) {
		case xml.StartElement:
			name := token.Name.Local
			fmt.Printf("Token name :%s\n", name)
			for _, attr := range token.Attr {
				attrName := attr.Name.Local
				attrValue := attr.Value
				fmt.Printf("an %s %s\n", attrName, attrValue)
			}
		case xml.EndElement:
			fmt.Println("end")
		case xml.CharData:
			content := string([]byte(token))
			fmt.Printf("this is %v\n", content)

		default:

		}
	}
}
/*
包中定义了若干 XML 标签类型：StartElement，Chardata（这是从开始标签到结束标签之间的实际文本），EndElement，Comment，Directive 或 ProcInst。

包中同样定义了一个结构解析器：NewParser 方法持有一个 io.Reader（这里具体类型是 strings.NewReader）并生成一个解析器类型的对象。还有一个 Token() 方法返回输入流里的下一个 XML token。在输入流的结尾处，会返回（nil，io.EOF）

XML 文本被循环处理直到 Token() 返回一个错误，因为已经到达文件尾部，再没有内容可供处理了。通过一个 type-switch 可以根据一些 XML 标签进一步处理。Chardata 中的内容只是一个 []byte，通过字符串转换让其变得可读性强一些。
 */