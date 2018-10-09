package main

import (
	"github.com/garyburd/redigo/redis"
	"github.com/astaxie/beego"
	"time"
	"fmt"
)

var (
	//定义常量
	RedisClient *redis.Pool
	REDIS_HOST  string
	REDIS_DB    int
)

func init() {
	// 从配置文件获取redis的ip以及db
	REDIS_HOST = beego.AppConfig.String("redis.host")
	REDIS_DB, _ = beego.AppConfig.Int("redis.db")
	// 建立连接池
	RedisClient = &redis.Pool{
		// 从配置文件获取maxidle以及maxactive，取不到则用后面的默认值
		MaxIdle:     beego.AppConfig.DefaultInt("redis.maxidle", 1),
		MaxActive:   beego.AppConfig.DefaultInt("redis.maxactive", 10),
		IdleTimeout: 180 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", REDIS_HOST)
			if err != nil {
				return nil, err
			}
			// 选择db
			c.Do("SELECT", REDIS_DB)
			return c, nil
		},
	}
}
//redis连接池配置完美版
func newpool() *redis.Pool {
	// 从配置文件获取redis的ip以及db
	REDIS_HOST = beego.AppConfig.String("redis.host")
	REDIS_DB, _ = beego.AppConfig.Int("redis.db")
	// 建立连接池
	RedisClient:=&redis.Pool{
		MaxActive:100,//最大连接数为100
		MaxIdle:10,//最大空闲连接数为10，也就是说，在不超过100连接的情况下要保持10连接，便于请过过来是迅速的反应
		IdleTimeout:180*time.Second,//空闲连接超过这个时间，会被关闭
		Wait:true,
		//用之前检查这个连接是不是健康的
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t)<time.Minute{
				return nil
			}
			_,err:=c.Do("PING")
			return err
		},
		Dial: func() (redis.Conn, error) {
			c,err:=redis.Dial("tcp",REDIS_HOST,
				redis.DialConnectTimeout(time.Duration(100)*time.Minute),
					redis.DialReadTimeout(time.Duration(100)*time.Minute),
						redis.DialWriteTimeout(time.Duration(100)*time.Minute),
			)
			if err!=nil{
				return nil,err
			}
			//选择数据库
			c.Do("SELECT",REDIS_DB)
			return c,nil
		},

	}
	return RedisClient
}


func main() {
	//获取操作数据库的连接
	rc := RedisClient.Get()
	defer rc.Close()
	//存数据
	key := "111"
	value := "222"
	//// 操作redis时调用Do方法，第一个参数传入操作名称（字符串），然后根据不同操作传入key、value、数字等
	// 返回2个参数，第一个为操作标识，成功则为1，失败则为0；第二个为错误信息
	n, err := rc.Do("SETNX", key, value)
	if err != nil {
		fmt.Println(err)
		return
	}
	if n == int64(1) {
		//设置过期时间
		n, _ := rc.Do("EXPIRE", key, 24*3600)
		if n == int64(1) {
			fmt.Println("success")
		}
	} else if n == int64(0) {
		fmt.Println("the key has already existed")
	}
	//取数据
	// 由于之前存的value是string类型，所以用redis.String将数据转换成string类型
	value1, err1 := redis.String(rc.Do("GET", key))
	if err1 != nil {
		fmt.Println("fail")
	}
	fmt.Println(value1)
	/*
	若value的类型为int，则用redis.Int转换
若value的类型为string，则用redis.String转换
若value的类型为json，则用redis.Byte转换
	 */

}
