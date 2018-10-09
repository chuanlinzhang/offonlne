package main

import (
	"net/url"
	"fmt"
	"encoding/json"
	"net/http"
	"io/ioutil"
)
const APPKEY = "*******************" //您申请的APPKEY
//1.查询
func Request1(){
	//请求地址
	juheURL :="http://api2.juheapi.com/video/query"

	//初始化参数
	param:=url.Values{}

	//配置请求参数,方法内部已处理urlencode问题,中文参数可以直接传参
	param.Set("key",APPKEY) //聚合key
	param.Set("kw","") //关键字
	param.Set("tag","") //分类
	param.Set("area","") //地区


	//发送请求
	data,err:=Get(juheURL,param)
	if err!=nil{
		fmt.Errorf("请求失败,错误信息:\\r\\n%v",err)
	}else{
		var netReturn map[string]interface{}
		json.Unmarshal(data,&netReturn)
		if netReturn["error_code"].(float64)==0{
			fmt.Printf("接口返回result字段是:\\r\\n%v",netReturn["result"])
		}
	}
}
// get 网络请求
func Get(apiURL string,params url.Values)(rs[]byte ,err error){
	var Url *url.URL
	Url,err=url.Parse(apiURL)
	if err!=nil{
		fmt.Printf("解析url错误:\\r\\n%v",err)
		return nil,err
	}
	//如果参数中有中文参数,这个方法会进行URLEncode
	Url.RawQuery=params.Encode()
	resp,err:=http.Get(Url.String())
	if err!=nil{
		fmt.Println("err:",err)
		return nil,err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}
func main() {
	Request1()
}