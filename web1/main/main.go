package main

import (
	"net/http"
	"fmt"
	"log"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println("path", r.URL.Path)
	fmt.Fprintf(w, "hello go")
}
/*
2.2 服务监听与请求处理过程

Go是通过一个ListenAndServe 监听服务，底层处理：初始化一个server对象，然后调用 net.Listen("tcp", addr)，监控我们设置的端口。

监控端口之后，调用 srv.Serve(net.Listener) 函数，处理接收客户端的请求信息。首先通过Listener 接收请求，其次创建一个Conn，最后单独开了一个goroutine，把这个请求的数据当做参数扔给这个conn去服务。go c.serve() 用户的每一次请求都是在一个新的goroutine去服务，相互不影响。

分配相应的函数处理请求: conn 首先会解析 request:c.readRequest(), 然后获取相应的handler:handler := c.server.Handler，这个是调用函数ListenAndServe 时候的第二个参数，例子传递的是nil，也就是为空，那么默认获取handler = DefaultServeMux。DefaultServeMux 是一个路由器，它用来匹配url跳转到其相应的handle函数

调用 http.HandleFunc("/", sayhelloName) 作用是注册了请求/的路由规则，将url 和handle 函数注册到DefaultServeMux 变量，最后调用DefaultServeMux 的ServeHTTP 方法，这个方法内部调用handle 函数。
 */
func main() {//ServeMux只是一个路由管理器

	http.HandleFunc("/6", sayhelloName)//sayhelloName是hander函数,handler处理器(函数):
	// 经过HandlerFunc结构包装的handler函数，它实现了ServeHTTP接口方法的函数。
	// 调用handler处理器的ServeHTTP方法时，即调用handler函数本身。
	err := http.ListenAndServe(":9090", nil)//
	if err != nil {
		log.Fatal("ListAndServer:", err)
	}
	//handler对象：实现了Handler接口ServeHTTP方法的结构。
	//任何结构体，只要实现了ServeHTTP方法，这个结构就可以称之为handler对象。
	// ServeMux会使用handler并调用其ServeHTTP方法处理请求并返回响应。
	//创建HTTP服务

	//创建一个http服务，大致需要经历两个过程，首先需要注册路由，即提供url模式和handler函数的映射，其次就是实例化一个server对象，并开启对客户端的监听
/*
例子传递的是nil，也就是为空，那么默认获取handler = DefaultServeMux。
DefaultServeMux 是一个路由器，它用来匹配url跳转到其相应的handle函数

调用 http.HandleFunc("/", sayhelloName) 作用是注册了请求/的路由规则，
将url 和handle 函数注册到DefaultServeMux 变量，最后调用DefaultServeMux 的ServeHTTP 方法，
这个方法内部调用handle 函数
 */
}
/*Web 服务代码实现
1.路由注册代码
   调用 http.HandleFunc(“/”, sayhelloName) 注册路由
   使用默认 ServeMux
   注册路由策略 DefaultServeMux
涉及数据结构

// :1900 ServeMux 默认实例是 DefaultServeMux
type ServeMux struct {
    mu    sync.RWMutex // 锁，由于请求涉及到并发处理，因此这里需要一个锁机制
    m     map[string]muxEntry // 路由规则，一个string对应一个mux实体，这里的string就是注册的路由表达式
    hosts bool // 是否在任意的规则中带有host信息
}

type muxEntry struct {
    explicit bool
    h        Handler // 路由处理器
    pattern  string  // url 匹配正则
}

type Handler interface {
    ServeHTTP(ResponseWriter, *Request)
}
服务监听代码
1 调用 err := http.ListenAndServe(“:9090”, nil) 监听端口
 */
 /*
 3.3 Go 代码的执行流程

调用Http.HandleFunc，按顺序做了几件事：

调用了DefaultServeMux的HandleFunc
调用了DefaultServeMux的Handle
往DefaultServeMux的map[string]muxEntry中增加对应的handler和路由规则
调用http.ListenAndServe(":9090", nil)，按顺序做了几件事情：

实例化Server
调用Server的ListenAndServe()
调用net.Listen(“tcp”, addr)监听端口
启动一个for循环，在循环体中Accept请求
对每个请求实例化一个Conn，并且开启一个goroutine为这个请求进行服务go c.serve()
读取每个请求的内容w, err := c.readRequest()
判断handler是否为空，如果没有设置handler（这个例子就没有设置handler），handler就设置为DefaultServeMux
调用handler的ServeHttp
在这个例子中，下面就进入到DefaultServeMux.ServeHttp
根据request选择handler，并且进入到这个handler的ServeHTTP mux.handler(r).ServeHTTP(w, r)
选择handler：
A 判断是否有路由能满足这个request（循环遍历ServerMux的muxEntry）
B 如果有路由满足，调用这个路由handler的ServeHttp
C 如果没有路由满足，调用NotFoundHandler的ServeHttp
  */