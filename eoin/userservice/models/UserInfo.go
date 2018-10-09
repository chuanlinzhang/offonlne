package models

import (
	"crypto/md5"
	"io"
	"time"
	"fmt"
	"encoding/json"
)

type UserInfo struct {
	Username string `json:"username" from:"username"` //用户名
	Password string `json:"password" from:"password"` //用户密码
	Sail string `json:"sail"`//MD5加密
	Info Profile //信息
}
type Profile struct {
	Email string
}



//md5加密的函数
func GetMD5(password ,s string)  string {
	h:=md5.New()
	_,err:=io.WriteString(h,password+s)
	if err!=nil{
		fmt.Println(err)
	}
	passmd5:=fmt.Sprintf("%x",h.Sum(nil))
	return passmd5
}
//用户注册函数
func UserRgister(username ,password string)  {
	t:=time.Now().Unix()
	tm:=time.Unix(t,0)
	sail:=tm.Format("2006-01-02 03:04:05 PM")
	md5:=GetMD5(password,sail)
	rc:=RedisClient.Get()
	defer rc.Close()

	ui:=&UserInfo{Username:username, Password:md5, Sail:sail, Info:Profile{Email:""}}
	val,err:=json.Marshal(ui)
	if err!=nil{
		fmt.Println(err)
	}
	rc.Do("hset",username)
}
