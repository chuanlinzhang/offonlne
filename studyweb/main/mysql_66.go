package main

import (
	"strconv"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"time"
	"log"

)

var db = &sql.DB{}
var err error

func init() {
	db, err = sql.Open("mysql", "root:123456@tcp(120.79.141.221:3306)/test2")
	fmt.Println(err, "init")

}

func main() {

	//insert()
	//query()
	//update()
	//query()
	//delete()


}




func Substr(str string, start int, end int) string {
	rs := []rune(str)
	length := len(rs)

	if start < 0 || start > length {
		return ""
	}

	if end < 0 || end > length {
		return ""
	}
	return string(rs[start:end])
}
func insert() {
	//方式1 insert
	//strconv,int 转string:strconv.Itoa(i)
	start := time.Now()
	for i := 1001; i <= 1100; i++ {
		//每次循环内部都会去连接池获取新的连接，效率低
		db.Exec("INSERT INTO user(uid,username,age) VALUES (?,?,?)", i, "user"+strconv.Itoa(i), i-1000, )

	}
	end := time.Now()
	fmt.Println("方式1 insert total time:", end.Sub(start).Seconds())
	//方式2 insert
	start = time.Now()
	for i := 1101; i <= 1200; i++ {
		//Prepare函数每次循环内部都会去连接池获取一个人新的连接，效率低
		stm, _ := db.Prepare("INSERT INTO user(uid,username,age) VALUES(?,?,?) ")
		stm.Exec(i, "user"+strconv.Itoa(i), i-1000)
		stm.Close()
	}
	end = time.Now()
	fmt.Println("方式2 insert total time:", end.Sub(start).Seconds())
	//方式3 insert////**************
	start = time.Now()
	stm, _ := db.Prepare("INSERT INTO user(Uid,username,age) VALUES (?,?,?)")
	//fmt.Println(err,"**********")
	for i := 1201; i <= 1300; i++ {
		//Exec内部并没有去获取连接。为什么效率还是低
		stm.Exec(i, "user"+strconv.Itoa(i), i-1000)
		//fmt.Println(err,"err insert3")
	}
	stm.Close()
	end = time.Now()
	fmt.Println("方式3 insert total time:", end.Sub(start).Seconds())
	//方式4 insert***************
	start = time.Now()
	//begin函数内部去获取连接
	tx, _ := db.Begin()
	for i := 1301; i <= 1400; i++ {
		//每次循环用的都是tx内部的连接，没有新建连接，效率高
		tx.Exec("INSERT INTO user(uid,username,age) VALUES (?,?,?)", i, "user"+strconv.Itoa(i), i-1000)

	}
	//最后释放tx内部的连接
	tx.Commit()
	end = time.Now()
	fmt.Println("方式4 insert total time：", end.Sub(start).Seconds())
	//方式5 insert
	start = time.Now()
	for i := 1401; i <= 1500; i++ {
		//begin函数每次循环内部都会去连接池获取一个新的连接，效率低
		tx, _ := db.Begin()
		tx.Exec("INSERT INTO user (uid,username,age) VALUES (?,?,?)", i, "user"+strconv.Itoa(i), i-1000)
		//Commit执行后链接也是释放
		tx.Commit()
	}
	end = time.Now()
	fmt.Println("方式5 insert total time：", end.Sub(start).Seconds())
}
func delete() {
	//方式1 delete
	start := time.Now()
	for i := 1001; i <= 1100; i++ {
		db.Exec("DELETE from user where uid=?", i)
	}
	end := time.Now()
	fmt.Println("方式1 delete total time:", end.Sub(start).Seconds())
	//方式2 delete
	start = time.Now()
	for i := 1101; i <= 1200; i++ {
		stm, _ := db.Prepare("DELETE FROM user WHERE uid=?")
		stm.Exec(i)
		stm.Close()
	}
	end = time.Now()
	fmt.Println("方式2 delete total time:", end.Sub(start).Seconds())
	//方式3 delete************
	start = time.Now()
	stm, _ := db.Prepare("DELETE FROM user WHERE uid=?")
	for i := 1201; i <= 1300; i++ {
		stm.Exec(i)
	}
	stm.Close()
	end = time.Now()
	fmt.Println("方式3 delete total time:", end.Sub(start).Seconds())
	//方式4 delete************
	start = time.Now()
	tx, _ := db.Begin()
	for i := 1301; i <= 1400; i++ {
		tx.Exec("DELETE FROM user WHERE uid=?", i)
	}
	tx.Commit()
	end = time.Now()
	fmt.Println("方式4 delete total time:", end.Sub(start).Seconds())
	//方式5 delete
	start = time.Now()
	for i := 1401; i <= 1500; i++ {
		tx, _ := db.Begin()
		tx.Exec("DELETE FROM  user WHERE uid=?", i)
		tx.Commit()
	}
	end = time.Now()
	fmt.Println("方式5 delete total time:", end.Sub(start).Seconds())

}

func update() {
	//方式1 update
	start := time.Now()
	for i := 1001; i <= 1100; i++ {
		db.Exec("update user SET age=? WHERE uid=?", i, i)
	}
	end := time.Now()
	fmt.Println("方式1 update total time:", end.Sub(start).Seconds())
	//方式2 update
	start = time.Now()
	for i := 1101; i <= 1200; i++ {
		stm, _ := db.Prepare("UPDATE user SET age=? WHERE uid=?")
		stm.Exec(i, i)
		stm.Close()
	}
	end = time.Now()
	fmt.Println("方式2 update total time:", end.Sub(start).Seconds())
	//方式3 update *********
	start = time.Now()
	stm, _ := db.Prepare("UPDATE user SET age=? WHERE uid=?")
	for i := 1201; i <= 1300; i++ {
		stm.Exec(i, i)
	}
	stm.Close()
	end = time.Now()
	fmt.Println("方式3 update total time:", end.Sub(start).Seconds())
	//方式4 update************
	start = time.Now()
	tx, _ := db.Begin()
	for i := 1301; i <= 1400; i++ {
		tx.Exec("UPDATE user SET age=? WHERE uid=?", i, i)
	}
	tx.Commit()
	end = time.Now()
	fmt.Println("方式4 update total time:", end.Sub(start).Seconds())
	//方式5 update
	start = time.Now()
	for i := 1401;i<=1500;i++{
		tx,_ := db.Begin()
		tx.Exec("UPdate user set age=? where uid=?",i,i)
		tx.Commit()
	}
	end = time.Now()
	fmt.Println("方式5 update total time:",end.Sub(start).Seconds())
}
func query() {
	//方式1 query   最快
	start := time.Now()
	rows, _ := db.Query("SELECT uid ,username from user ")
	defer rows.Close()
	for rows.Next() {
		var name string
		var id int
		if err := rows.Scan(&id, &name); err != nil {
			log.Fatal(err)
		}
	}
	end := time.Now()
	fmt.Println("方式1 query total time:", end.Sub(start).Seconds())
	//方式2 query
	start = time.Now()
	stm, _ := db.Prepare("SELECT uid,username FROM user ")
	defer stm.Close()
	rows, _ = stm.Query()
	for rows.Next() {
		var name string
		var id int
		if err := rows.Scan(&id, &name); err != nil {
			log.Fatal(err)
		}
	}
	end = time.Now()
	fmt.Println("方式2 query total time:", end.Sub(start).Seconds())
	//方式3 query
	start = time.Now()
	tx, _ := db.Begin()
	defer tx.Commit()
	rows, _ = tx.Query("SELECT uid,username FROM user ")
	defer rows.Close()
	for rows.Next() {
		var name string
		var id int
		if err := rows.Scan(&id, &name); err != nil {
			log.Fatal(err)
		}
	}
	end = time.Now()
	fmt.Println("方式3 query total time:", end.Sub(start).Seconds())
}
