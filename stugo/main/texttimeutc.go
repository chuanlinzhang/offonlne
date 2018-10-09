package main

import (

	"strconv"
	"fmt"
)

func main() {

	//local2, err2 := time.LoadLocation("Local")//服务器设置的时区
	//if err2 != nil {
	//	fmt.Println(err2)
	//}
	//loc:=time.Date(2018, 5, 15, 7, 10, 0, 0, local2)
	//s1:=loc.Format("2006-01-02 15:04:05")
	//fmt.Println(s1)
	//t2:=loc.AddDate(0,0,1)
	//s2:=t2.Format("2006-01-02 15:04:05")
	//fmt.Println(s2)
	//s3:=t2.In(time.UTC).Format("2006-01-02 15:04:05")
	//fmt.Println(s3)
	//app:=566.0/698.0*100
	//fmt.Println(app)
	//app1:=strconv.FormatFloat(app,'f',2,64)
	//fmt.Println(app1)
	cidlist:=make([]string,0)
	for i:=0;i<100000;i++  {
		ii:=strconv.Itoa(i)
		cidlist=append(cidlist,ii)
	}
	for i:=100;i<100001 ;i++  {
		ii:=strconv.Itoa(i)
		b:=containcid(ii,cidlist)
		if b==false{
			cidlist=append(cidlist,ii)
		}
	}
   fmt.Println(len(cidlist))
   //fmt.Println(len("62d3d700-c031-4843-9157-93469db8b885"))

}
func containcid(cid string, cidlist []string) bool {
	for _, c := range cidlist {
		if cid == c {
			return true
		}
	}
	return false
}