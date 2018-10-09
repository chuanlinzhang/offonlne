package tests
import (
	"testing"
	"net/http/httptest"
	"net/http"
	"fmt"
	"encoding/json"
)

var personResponse = []Person{
	{
		Name : "wahaha",
		Address : "shanghai",
		Age : 20,
	},
	{
		Name : "lebaishi",
		Address : "shanghai",
		Age : 10,
	},
}

var personResponseBytes, _ = json.Marshal(personResponse)

func TestPublishWrongResponseStatus(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write(personResponseBytes)
		if r.Method != "GET"{
			t.Errorf("Expected 'GET' request, got '%s'", r.Method)
		}
		if r.URL.EscapedPath() != "/person" {
			t.Errorf("Expected request to '/person', got '%s'", r.URL.EscapedPath())
		}
		r.ParseForm()
		topic := r.Form.Get("addr")
		if topic != "shanghai" {
			t.Errorf("Expected request to have 'addr=shanghai', got: '%s'", topic)
		}
	}))

	defer ts.Close()
	api := ts.URL
	fmt.Println("url:", api)
	resp, _ := GetInfo(api)

	fmt.Println("reps:", resp)
}
/*
解释一下：
>我们通过httptest.NewServer创建了一个测试的http server

>读请求设置通过变量r *http.Request，写变量（也就是返回值）通过w http.ResponseWriter

>通过ts.URL来获取请求的URL（一般都是<http://ip:port>）

>通过r.Method来获取请求的方法，来测试判断我们的请求方法是否正确

>获取请求路径：r.URL.EscapedPath()，本例中的请求路径就是"/person"

>获取请求参数：r.ParseForm，r.Form.Get("addr")

>设置返回的状态码：w.WriteHeader(http.StatusOK)

>设置返回的内容（这就是我们想要的结果）：w.Write(personResponseBytes)，注意w.Write()接收的参数是[]byte，因此需要将object对象列表通过json.Marshal(personResponse)转换成字节。



综上，我们可以通过不发送httptest来模拟出httpserver和返回值来进行自己代码的测试了
 */