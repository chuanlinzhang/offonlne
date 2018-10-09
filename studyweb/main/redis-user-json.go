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
	RedisClient2 *redis.Pool
	REDIS_HOST2  string
	REDIS_DB2    int
)

func init() {
	// 从配置文件获取redis的ip以及db
	REDIS_HOST2 = beego.AppConfig.String("redis.host")
	REDIS_DB2, _ = beego.AppConfig.Int("redis.db")
	// 建立连接池
	RedisClient2 = &redis.Pool{
		// 从配置文件获取maxidle以及maxactive，取不到则用后面的默认值
		MaxIdle:     beego.AppConfig.DefaultInt("redis.maxidle", 1),
		MaxActive:   beego.AppConfig.DefaultInt("redis.maxactive", 10),
		IdleTimeout: 180 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", REDIS_HOST2)
			if err != nil {
				return nil, err
			}
			// 选择db
			c.Do("SELECT", REDIS_DB2)
			return c, nil
		},
	}
}

type User struct {
	Name string
	Age  int
}

func main() {
	rc := RedisClient2.Get()
	defer rc.Close()
	key3 := "user"
	user := User{
		Name: "eoinzhang",
		Age:  18,
	}
	//将用户数据转化为json数据
	value, err := json.Marshal(user)
	if err != nil {
		fmt.Println(err)
	}
	n, err := rc.Do("SETNX", key3, value)
	if n == int64(1) {
		fmt.Println("success")
	}
	var users User
	value, err4 := redis.Bytes(rc.Do("GET", key3))
	if err4 != nil {
		fmt.Println(err4)
	}
	err5 := json.Unmarshal(value, &users)
	if err5 != nil {
		fmt.Println(err5)
	}

	fmt.Println(users)
}
