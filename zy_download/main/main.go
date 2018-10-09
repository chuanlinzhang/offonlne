package main

import (
	"regexp"
	"sync"
	"net/http"
	"io/ioutil"
	"os"
	"bufio"
	"io"
	"fmt"
	"strings"
	"time"
	"path"
)

//定义一些全局变量
var index_url string = "https://v.qq.com/"
/*
ptnIndexItem是一个个播放视频网页的链接
ptnVideoItem是为了匹配视频网页里的视频链接
dir 是你要下载的路径
 *///https://v.qq.com/x/cover/cltmsz5k25rt04w.html
//MustCompile 用来解析正则表达式 expr 是否合法，如果合法，则返回一个 Regexp 对象
var ptnIndexItem = regexp.MustCompile(`<a[^<>]+href *\= *[\"']?(\/[\d]+)\"[^<>]*title\=\"([^\"]*)\".*name.*>`)
var dir string = "./example_video"
var ptnVideoItem = regexp.MustCompile(`<a[^<>]+href *\= *[\"']?(https\:\/\/[^\"]+)\"[^<>]*download[^<>]*>`)
//增加一个等待组
var wg sync.WaitGroup
//写一个结构体，为了存放每一个视频的下载进度，为什么用int64是因为后面用到的文件大小和ContentLength都是int64
type DownList struct {
	Data map[string][]int64
	Lock sync.Mutex //（互斥锁）
}

//写一个检查错误的函数，用来检查每次可能出现错误时检查
func check(e error) {
	if e != nil {
		panic(e)
	}
}

//获取页面的方法、
/*
url：需要获取的网页
return：
    content 抓取到网页源码
    statusCode 返回的状态码
 */
func Get(url string) (content string, statusCode int) {
	resp, err1 := http.Get(url)
	if err1 != nil {
		statusCode = -100
		return
	}
	defer resp.Body.Close()
	data, err2 := ioutil.ReadAll(resp.Body)
	if err2 != nil {
		statusCode = -200
		return
	}
	statusCode = resp.StatusCode
	content = string(data)
	return
}

//用文件来记录这个视频的链接，每次下载之前先和这个文件里面的链接比较，有就不下载了，免得重复下载。
//得先创建一个你钟意的文件
/*
param:
filename:文件名
text：需要比较的字符串
return;
true 没
false 有
 */
func readOnLine(filename string, text string) bool {
	fi, err := os.OpenFile(filename, os.O_EXCL|os.O_APPEND, os.ModeAppend)
	check(err)
	defer fi.Close()
	text = text + "\n"
	br := bufio.NewReader(fi)
	for {
		a, c := br.ReadString('\n')
		if c == io.EOF {
			fmt.Println(text, "不存在，现在写")
			fi.WriteString(text)
			return true
		}
		if string(a) == text {
			fmt.Println("存在", text)
			break
		}
	}
	return false
}

/*
lock 为了防止后面goroutines 并发读写map时候报错 \033[2J
是用来清空屏幕的，做到类似top的效果，
虽然远不如top，但是能看到下载进度还是很爽的
 */
//输出下载进度的方法
func (downList *DownList) process() {
	for {
		downList.Lock.Lock() //加锁
		for key, arr := range downList.Data {
			fmt.Printf("%s progress:[%-50s] %d%% Done\n", key, strings.Repeat("#", int(arr[0]*50/arr[1])), arr[0]*100/arr[1])

		}
		downList.Lock.Unlock()
		time.Sleep(time.Second * 3)
		fmt.Printf("\033[2J")
	}
}

//下载视频的函数，
/*
url:视频链接
filename：本地文件名
downList：用来记录下载进度的一个结构体指针、
 */
func Down(url string, filename string, downList *DownList) bool {
	b := make([]byte, 1024)
	f, err := os.Create(filename)
	if err != nil {
		fmt.Println("创建文件失败")
		return false
	}
	defer f.Close()
	repo, err := http.Get(url)
	if err != nil {
		fmt.Println("获取资源失败")
		return false
	}
	defer repo.Body.Close()
	bufRead := bufio.NewReader(repo.Body)
	for {
		n, err := bufRead.Read(b)
		if err == io.EOF {
			break
		}
		f.Write(b[:n])
		fileInfo, err := os.Stat(filename)
		fileSize := fileInfo.Size()
		downList.Lock.Lock()
		downList.Data[filename] = []int64{fileSize, repo.ContentLength}
		downList.Lock.Unlock()
	}
	wg.Done()
	return true
}
func main() {
	//初始化downList 与 map
	var downListF DownList
	downListF.Data = make(map[string][]int64)
	downList := &downListF //取地址
	//首先获取index的网页的内容
	context, statusCode := Get(index_url)

	if statusCode != 200 {
		fmt.Println("error")
		return
	}
	/*    提取并复制到二维数组
	html_result              [][]string
	html_result[n]          []string    匹配到的链接
	html_result[n][0]          string     全匹配数据
	html_result[n][1]         url        string
	html_result[n][2]         title        string
*/
	html_result := ptnIndexItem.FindAllStringSubmatch(context, -1)
	length := len(html_result)
	fmt.Println(context)
	fmt.Println(length)
	go downList.process() //输出下载进度的方法
	for i := 0; i < length; i++ {
		v := html_result[i]
		video_html, video_status := Get(index_url + v[1])
		if video_status != 200 {
			fmt.Println("error")
			continue
		}
		video_result := ptnVideoItem.FindAllStringSubmatch(video_html, -1)
		ok := readOnLine("test", v[1])
		if len(video_result) > 0 && len(video_result[0]) > 0 && ok {
			//fmt.Println(video_result[0][1])
			wg.Add(1)
			dirFile := path.Join(dir, html_result[i][2])
			go Down(video_result[0][1], dirFile, downList)

		}

	}
	wg.Wait()
	return
}
