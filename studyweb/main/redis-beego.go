package main

import (
	"github.com/astaxie/beego/cache"
	_ "github.com/astaxie/beego/cache/redis"
	"log"
	"time"
	"fmt"
)

func main() {
	redis, err := cache.NewCache("redis", `{"conn":"120.79.141.221:6379"}`)
	if err != nil {
		log.Fatal(err)
	}
	err1 := redis.Put("name", "eoinzhang", time.Duration(time.Second*1))
	if err1 != nil {
		log.Fatal(err1)
	}
	fmt.Println(string(redis.Get("name").([]uint8)))//eoinzhang
}
