package main

import (
	"strings"
	"fmt"
)

func main() {
	// Print 将参数列表 a 中的各个参数转换为字符串并写入到标准输出中。
	// 非字符串参数之间会添加空格，返回写入的字节数
	//fmt.Print("fmt")
	//fmt.Println(1615154, 56664)
	// Printf 将参数列表 a 填写到格式字符串 format 的占位符中。
	// 填写后的结果写入到标准输出中，返回写入的字节数。
	//fmt.Printf("%d:", 61161)

	// 功能同上面三个函数，只不过将转换结果写入到 w 中。
	//func Fprint(w io.Writer, a ...interface{}) (n int, err error)
	//func Fprintln(w io.Writer, a ...interface{}) (n int, err error)
	//func Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error)

	// 功能同上面三个函数，只不过将转换结果以字符串形式返回。
	//func Sprint(a ...interface{}) string
	//func Sprintln(a ...interface{}) string
	//func Sprintf(format string, a ...interface{}) string

	// 功能同 Sprintf，只不过结果字符串被包装成了 error 类型。
	//func Errorf(format string, a ...interface{}) error

	// Scan 从标准输入中读取数据，并将数据用空白分割并解析后存入 a 提供
	// 的变量中（换行符会被当作空白处理），变量必须以指针传入。
	// 当读到 EOF 或所有变量都填写完毕则停止扫描。
	// 返回成功解析的参数数量。
	//func Scan(a ...interface{}) (n int, err error)

	// Scanln 和 Scan 类似，只不过遇到换行符就停止扫描。
	//func Scanln(a ...interface{}) (n int, err error)

	// Scanf 从标准输入中读取数据，并根据格式字符串 format 对数据进行解析，
	// 将解析结果存入参数 a 所提供的变量中，变量必须以指针传入。
	// 输入端的换行符必须和 format 中的换行符相对应（如果格式字符串中有换行
	// 符，则输入端必须输入相应的换行符）。
	// 占位符 %c 总是匹配下一个字符，包括空白，比如空格符、制表符、换行符。
	// 返回成功解析的参数数量。
	//func Scanf(format string, a ...interface{}) (n int, err error)

	// 功能同上面三个函数，只不过从 r 中读取数据。
	////func Fscan(r io.Reader, a ...interface{}) (n int, err error)
	//func Fscanln(r io.Reader, a ...interface{}) (n int, err error)
	//func Fscanf(r io.Reader, format string, a ...interface{}) (n int, err error)

	// 功能同上面三个函数，只不过从 str 中读取数据。
	//func Sscan(str string, a ...interface{}) (n int, err error)
	//func Sscanln(str string, a ...interface{}) (n int, err error)
	//func Sscanf(str string, format string, a ...interface{}) (n int, err error)

	// 示例

	// 对于 Scan 而言，回车视为空白
	//func main() {
	//	a, b, c := "", 0, false
	//	fmt.Scan(&a, &b, &c)
	//	fmt.Println(a, b, c)
	//	// 在终端执行后，输入 abc 1 回车 true 回车
	//	// 结果 abc 1 true
	//}

	// 对于 Scanln 而言，回车结束扫描
	//func main() {
	//	a, b, c := "", 0, false
	//	fmt.Scanln(&a, &b, &c)
	//	fmt.Println(a, b, c)
	//	// 在终端执行后，输入 abc 1 true 回车
	//	// 结果 abc 1 true
	//}

	// 格式字符串可以指定宽度
	//func main() {
	//	a, b, c := "", 0, false
	//	fmt.Scanf("%4s%d%t", &a, &b, &c)
	//	fmt.Println(a, b, c)
	//	// 在终端执行后，输入 1234567true 回车
	//	// 结果 1234 567 true
	//}
//替换原字符串的中 部分字符串。n为替换的个数
	str := strings.Replace("12345666", "6", "2", 3)
	fmt.Println(str)
}
