package main

import (

	"fmt"
	"github.com/astaxie/beego"


)
type Functions struct {
	Function	string		`json:"function"`
}

func main() {
	//env:=os.Getenv("ENV_CLUSTER")
    err:= beego.LoadAppConfig("ini","./conf/app2.conf")
	//beego.AppPath="./conf/app2.conf"
fmt.Println(err)
	fmt.Println(beego.AppConfig.String("app"))//2
	//fmt.Println(conf)
	r:=Functions{"1"}
	s:=[]int32(r.Function)
	fmt.Println(s)
}
