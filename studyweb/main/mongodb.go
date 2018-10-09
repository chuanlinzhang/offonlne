package main

import (
	"gopkg.in/mgo.v2"
	_ "gopkg.in/mgo.v2/bson"
	"fmt"
)

type Person struct {
	Name  string `bson:"name"`
	Phone string `bson:"phone"`
}
//注意User的字段首字母大写，不然不可见。通过bson:”name”这种方式可以定义MongoDB中集合的字段名
//如果不定义，mgo自动把struct的字段名首字母小写作为集合的字段名。如果不需要获得id_，Id_可以不定义，在插入的时候会自动生成
func main() {
	//通过方法Dial()来和MongoDB服务器建立连接
	session, err := mgo.Dial("mongodb://localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	/*
	如果不在本机或端口不同，传入相应的地址即可。如：

mongodb://myuser:mypass@localhost:40001,otherhost:40001/mydb
通过Session.DB()来切换相应的数据库。

func (s *Session) DB(name string) *Database
如切换到test数据库。

db := session.DB("test")
通过Database.C()方法切换集合（Collection），这样我们就可以通过对集合进行增删查改操作了。

func (db *Database) C(name string) *Collection
如切换到`users`集合。

c := db.C("users")
	 */
	//Optional.Switchthesessiontoamonotonicbehavior.
	session.SetMode(mgo.Monotonic, true)
	c := session.DB("test").C("people")
	//插入数据
	err = c.Insert(&Person{"Ale", "+555381169639"},
		&Person{"Cla", "+555384028510"})
	if err != nil {
		panic(err)
	}
	//查询所有
	var result []Person//定义一个切片
	c.Find(nil).All(&result)
	for _,v:=range result{
		fmt.Println(v.Name,v.Phone)
	}
   //有条件查询
	//result := Person{}
	//err = c.Find(bson.M{"phone": "+555381169639"}).One(&result)
	//if err != nil {
	//	panic(err)
	//}
	///*
	//"phone":这里必须是db中的字段名，
	// */
	//
	//fmt.Println("Phone:",result.Name, result.Phone)
	/*
	单条件查询

=($eq)
c.Find(bson.M{"name": "Jimmy Kuu"}).All(&users)
!=($ne)
c.Find(bson.M{"name": bson.M{"$ne": "Jimmy Kuu"}}).All(&users)
>($gt)
c.Find(bson.M{"age": bson.M{"$gt": 32}}).All(&users)
<($lt)
c.Find(bson.M{"age": bson.M{"$lt": 32}}).All(&users)
>=($gte)
c.Find(bson.M{"age": bson.M{"$gte": 33}}).All(&users)
<=($lte)
c.Find(bson.M{"age": bson.M{"$lte": 31}}).All(&users)
in($in)
c.Find(bson.M{"name": bson.M{"$in": []string{"Jimmy Kuu", "Tracy Yu"}}}).All(&users)
多条件查询

and($and)
c.Find(bson.M{"name": "Jimmy Kuu", "age": 33}).All(&users)
or($or)
c.Find(bson.M{"$or": []bson.M{bson.M{"name": "Jimmy Kuu"}, bson.M{"age": 31}}}).All(&users)

	 */
	 //修改
	 /*

通过func (*Collection) Update来进行修改操作。

func (c *Collection) Update(selector interface{}, change interface{}) error
注意修改单个或多个字段需要通过$set操作符号，否则集合会被替换。

修改字段的值($set)

c.Update(bson.M{"_id": bson.ObjectIdHex("5204af979955496907000001")},
    bson.M{"$set": bson.M{
        "name": "Jimmy Gu",
        "age":  34,
    }})
inc($inc)

字段增加值

c.Update(bson.M{"_id": bson.ObjectIdHex("5204af979955496907000001")},
    bson.M{"$inc": bson.M{
        "age": -1,
    }})
push($push)

从数组中增加一个元素

c.Update(bson.M{"_id": bson.ObjectIdHex("5204af979955496907000001")},
    bson.M{"$push": bson.M{
        "interests": "Golang",
    }})
pull($pull)

从数组中删除一个元素

c.Update(bson.M{"_id": bson.ObjectIdHex("5204af979955496907000001")},
    bson.M{"$pull": bson.M{
        "interests": "Golang",
    }})
删除

c.Remove(bson.M{"name": "Jimmy Kuu"})
	  */
}
