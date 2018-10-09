package main

import (
	"encoding/json"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"github.com/astaxie/beego"
	"time"
)
var (
	//定义常量
	RedisClient1 *redis.Pool
	REDIS_HOST1  string
	REDIS_DB1   int
)

func init() {
	// 从配置文件获取redis的ip以及db
	REDIS_HOST1 = beego.AppConfig.String("redis.host")
	REDIS_DB1, _ = beego.AppConfig.Int("redis.db")
	// 建立连接池
	RedisClient1 = &redis.Pool{
		// 从配置文件获取maxidle以及maxactive，取不到则用后面的默认值
		MaxIdle:     beego.AppConfig.DefaultInt("redis.maxidle", 1),
		MaxActive:   beego.AppConfig.DefaultInt("redis.maxactive", 10),
		IdleTimeout: 180 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", REDIS_HOST1)
			if err != nil {
				return nil, err
			}
			// 选择db
			c.Do("SELECT", REDIS_DB1)
			return c, nil
		},
	}
}
func main() {
	//获取操作数据库的连接
	rc := RedisClient1.Get()
	defer rc.Close()
	//存json数据
	key := "jj"
	imap1 := map[string]string{"key1": "666", "key2": "666"}
	//将map转化为json数据
	value1, _ := json.Marshal(imap1)
	//存入redis
	n, err := rc.Do("SETNX", key, value1)
	if err != nil {
		fmt.Println(err)
	}
	if n == int64(1) {
		fmt.Println("success") //存入成功
	}
	//取json数据
	//先声明imap2用来装数据
	var imap2 map[string]string
	key2 := "jj"
	//json数据在go中是[]byte类型的。所以用redis.Bytes转换
	value2, err2 := redis.Bytes(rc.Do("GET", key2))
	if err2 != nil {
		fmt.Println(err2)
	}
	//将json数据解析成map数据
	err3 := json.Unmarshal(value2, &imap2)
	if err3 != nil {
		fmt.Println(err3)
	}
	fmt.Println(imap2["key1"])
	fmt.Println(imap2["key2"])
}
