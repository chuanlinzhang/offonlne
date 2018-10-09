package main

import (
	"net"
	"github.com/astaxie/beego/logs"
	"strings"
	"bufio"
	"os"
)

func main() {
	conn,err:=net.Dial("tcp","localhost:1208")
	if err!=nil{
		logs.Error(err)
		return
	}
	connHandler(conn)
}
func connHandler(c net.Conn)()  {
	defer c.Close()
	reader:=bufio.NewReader(os.Stdin)
	buf:=make([]byte,1024)
	for {
		input,err:=reader.ReadString('\n')
		if err!=nil{
			logs.Error(err)
			break
		}
		input=strings.TrimSpace(input)
		if input=="quit"{
			return
		}
		c.Write([]byte(input))
		cnt,err:=c.Read(buf)
		if err!=nil{
			logs.Error(err)
			continue
		}
		logs.Info(string(buf[0:cnt]))
	}
}