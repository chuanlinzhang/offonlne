package main

import (
	"time"

	"github.com/garyburd/redigo/redis"
	"github.com/astaxie/beego"
	"encoding/json"
	"fmt"
	"log"
	"gopkg.in/mgo.v2"

	"strconv"
	"regexp"
	"gopkg.in/mgo.v2/bson"

)

type stu struct {
	Id int `bson:"id"`
	Name string `bson:"name"`
	Password string `bson:"password"`

}
type user struct {
	Name string
	Password string
}

var (
	//定义常量
	RedisClient3 *redis.Pool
	REDIS_HOST3  string
	REDIS_DB3    int
)
var collectionName *mgo.Collection
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
	url := beego.AppConfig.String("mongodb::dburl")
	Dbname := beego.AppConfig.String("mongodb::dbname")
	session, err := mgo.Dial(url)
	if err != nil {
		log.Fatal("数据库连接失败")
		return
	}
	conn :=session.Copy()
	DB := conn.DB(Dbname)
	collectionName = DB.C("w:u:")
}
func main() {
	////redis写
	rc := RedisClient3.Get()
	defer rc.Close()
	user1 := stu{
		Id:3,
		Name: "eoinzhang",
		Password:  "666",
	}
	//val,err:=json.Marshal(user1)
	//if err!=nil{
	//	fmt.Println(err)
	//}
	userKey:="w:u:"+strconv.Itoa(user1.Id)
	//_, err = rc.Do("hset", "w:u:", userKey, val)
	//if err != nil {
	//	fmt.Println(err)
	//}
	////mongodb写
	//err=collectionName.Insert(&user1)
	//if err!=nil{
	//	fmt.Println(err)
	//}
	//
	result, err := rc.Do("hget", "w:u:", userKey)

	if err!=nil{
		fmt.Println(err)
	}

	val,err:=redis.Bytes(result,err)
	if err!=nil{
		fmt.Println(err)
	}
	var u stu
	err=json.Unmarshal(val,&u)
	if err!=nil{
		fmt.Println(err)
	}
	var valid = regexp.MustCompile("[0-9]")
	id:=valid.FindAllStringSubmatch(userKey, -1)[0][0]
	fmt.Println(id)
	i,_:=strconv.Atoi(id)
	var user11  []stu
	err=collectionName.Find(bson.M{"id": i}).All(&user11)
	if err!=nil{
		fmt.Println(err)
	}
	if u==user11[0]{
		fmt.Println("统一")
	}

}
