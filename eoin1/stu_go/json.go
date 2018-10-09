package main

import (
	"encoding/json"
	"fmt"
)

type Server struct {
	ServerName string
	ServerIp string
}
type Serverslice struct {
	Servers []Server
}

func main()  {
	//解析一个json
	var s Serverslice
	str:=`{"servers":[{"serverName":"GZZ_VPN","serverIp":"127.0.0.3"},{"serverName":"SZ_VPN","serverIp":"127.0.0.4"}]}`
	//这里是json_decode
	json.Unmarshal([]byte(str),&s)
	fmt.Println(s)
	//添加一个json
	s.Servers=append(s.Servers,Server{ServerName:"123",ServerIp:"127.0.0.5"})
	s.Servers=append(s.Servers,Server{ServerName:"456",ServerIp:"127.0.0.6"})

	b,err:=json.Marshal(s)
	if err!=nil{
		fmt.Println(err)
	}
	fmt.Println(string(b))
}