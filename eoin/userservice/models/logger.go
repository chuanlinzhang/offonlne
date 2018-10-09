package models

import (
	"gopkg.in/mgo.v2"
	"encoding/json"
	"github.com/astaxie/beego"
	"time"
	"fmt"
)

//日志
//操作mongodb数据库的记录器
type MongoDBLogger struct {
	session      *mgo.Session `json:"-"`          //数据库连接时用到的session
	Url          string       `json:"url"`        //连接数据库的路径
	DbName       string       `json:"db_name"`    //数据库名
	PreColleName string       `json:"colle_name"` //集合名称
}

//日志信息
type lOgMsg struct {
	Msg string `json:"msg" bson:"msg"`
}

/*
实现beego框架中Logger接口，重写了4中方法
 */
//初始化记录器方法
func (this *MongoDBLogger) Init(config string) error { //config在设置使用本记录器是传入的参数
	if config != "" { //如果参数不为空
		if err := json.Unmarshal([]byte(config), this); err != nil { //用json解析
			return err
		}
	}
	if this.Url == "" {
		this.Url = beego.AppConfig.String("mongodb::url")
	}
	if this.DbName == "" {
		this.DbName = beego.AppConfig.String("mongodb:dbName")
	}
	if this.PreColleName == "" {
		this.PreColleName = beego.AppConfig.String("log::preLogFileName")
	}
	//初始化数据库
	s, err := mgo.Dial(this.Url)
	if err != nil {
		return err
	}
	s.SetMode(mgo.Monotonic, true)
	this.session = s //复制给session
	return nil
}

//把产生的日志信息写如数据库//msg日志信息beego会产生
func (this *MongoDBLogger) WriteMsg(when time.Time, msg string, level int) error {
	if this.session == nil {
		return fmt.Errorf("error not connect to host")
	}
	//获得操作数据库的连接
	conn := this.session.Copy()
	defer conn.Close()
	c := conn.DB(this.DbName).C(this.PreColleName + time.Now().Format("_2006_01_02"))
	//插入日志
	err := c.Insert(&lOgMsg{msg})
	return err
}

//摧毁连接的方法
func (this *MongoDBLogger) Destroy() {
	if this.session != nil {
		this.session.Close()
	}
}

func (this *MongoDBLogger) Flush() {

}
