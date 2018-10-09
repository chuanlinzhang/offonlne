package main

import (
	"gopkg.in/mgo.v2"
	"fmt"
	"gopkg.in/mgo.v2/bson"
)

type Usertest struct {
	Name      string   `bson:"name"`
	Age       int      `bson:"age"`
	Interests []string `bson:"interests"`
}

func main() {
	//mongodb://myuser:mypass@localhost:40001,otherhost:40001/mydb
	session, err := mgo.Dial("mongodb://localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	c := session.DB("test").C("peopletest")

	//err = c.Insert(&Usertest{"eoin", 20, []string{"11", "11", "11"}},
	//	&Usertest{"eoin1", 20, []string{"22", "22", "22"}})
	//if err != nil {
	//	fmt.Println("插入错误")
	//}
	//查询所以
	var uu []Usertest
	var uu1 Usertest
	c.Find(nil).All(&uu)
	fmt.Println(uu)
	// 使用迭代器，可以避免一次占用较大内存
	iter := c.Find(nil).Iter()
	for iter.Next(&uu1) {
		fmt.Println(uu1,"********")
	}
	/*
	// 用bson.M结构接收，当你不了解返回的数据结构格式时，可以用这个先查看，然后再定义struct格式
    // 在处理mongodb组合查询时，经常这么干
    result := bson.M{}
    err = c.Find(nil).One(&result)
    if err != nil {
        panic(err)
    }
    fmt.Println(result)
	 */

	 /*
	  // 查找表总数
    count, err := c.Count()

    // 结合find条件查找
    count, err = c.Find(bson.M{"name": "Tom"}).Count()
	 // 按照age字段降序排列，如果升序去掉横线"-"就可以了
    err := c.Find(nil).Sort("-age").All(&users)
	  // 表示从偏移位置为2的地方开始取两条记录  skip 的效率低
    err := c.Find(nil).Sort("-age").Skip(2).Limit(2).All(&users)
	  */
	//条件查询
	c.Find(bson.M{"name": "eoin"}).All(&uu)
	fmt.Println(uu)
	err = c.Find(bson.M{"age": bson.M{"$gt": 10}}).All(&uu)
	fmt.Println(err, uu)
	//查询，name的值至少为其中之一"eoin","eoin1"
	c.Find(bson.M{"name": bson.M{"$in": []string{"eoin", "eoin1"}}}).All(&uu)
	fmt.Println(uu)
	//是否包含这个键($exists)
	c.Find(bson.M{"name": bson.M{"$exists": true}}).All(&uu)
	fmt.Println(uu)
	//查询键值为null的字段
	c.Find(bson.M{"name": bson.M{"$in": []interface{}{nil}, "$exists": true}}).All(&uu)
	fmt.Println(uu)
	//模糊查询($regex)
	c.Find(bson.M{"age": bson.M{"$regex": "^[0-9]+"}}).All(&uu)
	fmt.Println(uu,"regex")
	//查询键值为长度是size(值为数组)的数据
	c.Find(bson.M{"interests": bson.M{"$size": 3}}).All(&uu)
	fmt.Println(uu)
	//查询数组中包含所有值（进行匹配，满足输出）（部分顺序，只看包含与否）
	c.Find(bson.M{"interests": bson.M{"$all": []string{"11", "11", "11"}}}).All(&uu)
	fmt.Println(uu)
	//如果数组只有一项内容
	c.Find(bson.M{"interests": "11"}).All(&uu)
	fmt.Println(uu)
	//如果要查询数组指定位置,查询Interests的第二个元素为"33"的所有人
	c.Find(bson.M{"interests.2": "22"}).All(&uu)
	fmt.Println(uu)
	//and($and)
	c.Find(bson.M{"name": "eoin", "age": 20}).All(&uu)
	fmt.Println(uu)
	/*
	conditions := []bson.M{
    bson.M{"content.type": TypeTopic},
    bson.M{"content.markdown": bson.M{"$regex": bson.RegEx{"go", "i"}}},
    bson.M{"content.markdown": bson.M{"$regex": bson.RegEx{"python", "i"}}},
}
c.Find(bson.M{"$and": conditions})
	 */
	//or($or)
	c.Find(bson.M{"$or": []bson.M{bson.M{"name": "eoin"}, bson.M{"age": 20}}}).All(&uu)
	fmt.Println(uu,"_________")
	/////修改
	//2.1、($set)
	//修改字段的值,前面是修改的条件，后面是修改
	c.Update(bson.M{"name": "eoin"}, bson.M{"$set": bson.M{"age": 22}})
	//2.2、inc($inc) 修改字段的值,前面是修改的条件，后面是“变化的修改”
	c.Update(bson.M{"name": "eoin"}, bson.M{"$inc": bson.M{"age": +1}})
	//push($push)，当某个字段的值为一个数组时 ，可以进行数组元素的增加(这就向interests数组中增加了golang元素)
	c.Update(bson.M{"name": "eoin"}, bson.M{"$push": bson.M{"interests": "golang"}})
	//pull($pull) ，同理删除元素
	c.Update(bson.M{"name": "eoin"}, bson.M{"$pull": bson.M{"interests": "golang"}})

	//删除数据文档
	//c.Remove(bson.M{"name": "eoin"})

	//如果想改变数组中元素里面的内容，只有把原来的元素从数组删除，然后再增加修改之后的新的元素进入数组
}
