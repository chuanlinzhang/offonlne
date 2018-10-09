package main

import (
	"net/http"
	"fmt"
	"strings"
	"html/template"
	"log"
	"net/url"
)

func sayhelloName2(res http.ResponseWriter, req *http.Request) {
	//解析url传递的参数，对于POST则解析响应包的主体（request body）
	req.ParseForm() //解析form表单
	//注意:如果没有调用ParseForm方法，下面无法获取表单的数据
	fmt.Println(req.Form) //这些信息是输出到服务器端的打印信息
	fmt.Println("path", req.URL.Path)
	fmt.Println("scheme", req.URL.Scheme)
	fmt.Println(req.Form["url_long"])
	for k, v := range req.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(res, "hello astaxie")
}
func login(res http.ResponseWriter, req *http.Request) {
	fmt.Println("method:", req.Method) //获取请求的方法
	if req.Method == "GET" {
		t, _ := template.ParseFiles("main/login.gtpl") //当需要加载静态文件是最前面不需要加/
		t.Execute(res, nil)
	} else {
		req.ParseForm() //解析form表单
		fmt.Println(req.Form)
		//请求的是登陆数据，那么执行登陆的逻辑判断
		fmt.Println("username:", req.Form["username"])
		fmt.Println("password:", req.Form["password"])
		fmt.Println(req.Form.Get("username"))//因为如果字段不存在，通过该方式获取的是空值。 但是通过r.Form.Get()只能获取单个的值
		v := url.Values{} //request.Form是一个url.Values类型，里面存储的是对应的类似key=value的信息
		v.Set("name", "Ava")
		v.Add("friend", "Jess")
		v.Add("friend", "Sarah")
		v.Add("friend", "Zoe")
		// v.Encode() == "name=Ava&friend=Jess&friend=Sarah&friend=Zoe"

		// 打印出: Ava
		fmt.Println(v.Get("name"))
		// 打印出: Jess
		fmt.Println(v.Get("friend"))
		// 打印出: [Jess Sarah Zoe]
		fmt.Println(v["friend"])

		// 注意, 当使用多个值时, 请使用Add. 如果使用Set则后面的值会覆盖前面的值
	}
}
func main() {
	http.HandleFunc("/", sayhelloName2)
	http.HandleFunc("/login", login) //当写访问路径是需要加入/
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe", err)
	}
}

/*
Tips:
Request本身也提供了FormValue()函数来获取用户提交的参数。 如req.Form["username"]也可写成req.FormValue("username")。
调用req.FormValue时会自动调用req.ParseForm，
所以不必提前调用。 req.FormValue只会返回同名参数中的第一个，若参数不存在则返回空字符串
 */
/*
如果我们是判断正整数，那么我们先转化成int类型，然后进行处理

getint, err := strconv.Atoi(req.Form.Get("age"))

if err != nil {
    //数字转化出错了，那么可能就不是数字
}

//接下来就可以判断这个数字的大小范围了
if getint > 100 {
    //太大了
}
 */
 //里面包括很多路由规则（ routes ）、服务接口（ services ）和处理函数（ handlers ）等