package main

import (
	"github.com/garyburd/redigo/redis"
	"github.com/astaxie/beego"
	"time"
	"fmt"
	_ "reflect"


	"encoding/json"
)

var (
	//定义常量
	RedisClient3 *redis.Pool
	REDIS_HOST3  string
	REDIS_DB3    int
)

func init() {
	// 从配置文件获取redis的ip以及db
	REDIS_HOST3 = beego.AppConfig.String("redis.host")
	REDIS_DB3, _ = beego.AppConfig.Int("redis.db")
	// 建立连接池
	RedisClient3 = &redis.Pool{
		// 从配置文件获取maxidle以及maxactive，取不到则用后面的默认值
		MaxIdle:     beego.AppConfig.DefaultInt("redis.maxidle", 1),
		MaxActive:   beego.AppConfig.DefaultInt("redis.maxactive", 10),
		IdleTimeout: 180 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", REDIS_HOST3)
			if err != nil {
				return nil, err
			}
			// 选择db
			c.Do("SELECT", REDIS_DB3)
			return c, nil
		},
	}
}

type User1 struct {
	Name string
	Age  int
}
/*
hset(key, field, value)：向名称为key的hash中添加元素field,value
hget(key, field)：返回名称为key的hash中field对应的value
hmget(key, (fields))：返回名称为key的hash中field i对应的value
hmset(key, (fields,values))：向名称为key的hash中添加元素field
hincrby(key, field, integer)：将名称为key的hash中field的value增加integer
hexists(key, field)：名称为key的hash中是否存在键为field的域
hdel(key, field)：删除名称为key的hash中键为field的域
hlen(key)：返回名称为key的hash中元素个数
hkeys(key)：返回名称为key的hash中所有键
hvals(key)：返回名称为key的hash中所有键对应的value
hgetall(key)：返回名称为key的hash中所有的键（field）及其对应的value
 */


func main() {
	rc := RedisClient3.Get()
	defer rc.Close()

	user1 := User1{
		Name: "eoinzhang",
		Age:  18,
	}
	val,err:=json.Marshal(user1)
	if err!=nil{
		fmt.Println(err)
	}
	_, err = rc.Do("hset", "myhash6", "user1", val)
	if err != nil {
		fmt.Println(err)
	}
	result, err := rc.Do("hget", "myhash6", "user1")
	//fmt.Println(reflect.TypeOf(result)) //[]uint8
	if err != nil {
		fmt.Println(err)
	} else {
		var u User1
		val,_:=redis.Bytes(result,err)
		err:=json.Unmarshal(val,&u)
		if err!=nil{
			fmt.Println(err)
		}
		fmt.Println(u)
		//fmt.Println(string(result.([]uint8)))
		//fmt.Println(err)
		fmt.Println(u.Name,u.Age) //mobike
	}

	//user2 := User1{
	//	Name: "eoin",
	//	Age:  20,
	//}
	//
	////hset//"hset"命令，myhash存入的hash表   bike1其中条数据的键 mobike键所对应的值
	//_, err := rc.Do("hset", "myhash2", "bike1", "mobike")
	//if err != nil {
	//	fmt.Println(err)
	//}
	////hget
	//result, err := rc.Do("hget", "myhash2", "bike1")
	//fmt.Println(reflect.TypeOf(result)) //[]uint8
	//if err != nil {
	//	fmt.Println(err)
	//} else {
	//	fmt.Println(string(result.([]uint8))) //mobike
	//}
	//
	////hmset/hmget
	//_, err = rc.Do("hmset", "myhash2", "bike2", "bluegogo", "bike3", "xiaoming", "bike4", "xiaolan")
	//if err != nil {
	//	fmt.Println(err)
	//}else{
	//	value,err:=redis.Values(rc.Do("hmget", "myhash2", "bike1", "bike2", "bike3", "bike4"))
	//  if err!=nil{
	//  	fmt.Println(err)
	//  }else {
	//  	for _,v:=range value{
	//  		fmt.Println(string(v.([]byte)))
	//	}
	//	fmt.Println()
	//  }
	//	}
		//hgetall
		//var key, val string
		//result1,err:=redis.Values(rc.Do("hgetall","myhash1"))
		//if err!=nil{
		//	fmt.Println(err)
		//}else {
		//	for i,re:=range result1{
		//		k1:=string(re.([]uint8))
		//		fmt.Println("k1:"+k1)
		//		if i%2==0{
		//			key=k1
		//			continue
		//		}else {
		//			val=k1
		//		}
		//		fmt.Println(key,val)
		//	}
		//}


}
