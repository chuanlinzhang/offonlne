package main


import "database/sql"
import (
	_ "github.com/go-sql-driver/mysql"
	"fmt"
)
/*
CREATE TABLE `userinfo` (
	`uid` INT(10) NOT NULL AUTO_INCREMENT,
	`username` VARCHAR(64) NULL DEFAULT NULL,
	`departname` VARCHAR(64) NULL DEFAULT NULL,
	`created` DATE NULL DEFAULT NULL,
	PRIMARY KEY (`uid`)
);

CREATE TABLE `userdetail` (
	`uid` INT(10) NOT NULL DEFAULT '0',
	`intro` TEXT NULL,
	`profile` TEXT NULL,
	PRIMARY KEY (`uid`)
)
 */
func main() {
	//db, err := sql.Open("mysql", "root:123456@tcp(120.79.141.221:3306)/test1")//test1为数据库名
	db, err := sql.Open("mysql", "root:password@tcp(192.168.103.88:3306)/test1")//test1为数据库名
	fmt.Println(db)
	chechErr(err)
	//插入数据
	stmt, err := db.Prepare("INSERT userinfo SET username=?,departname=?,created=?")
	chechErr(err)
	res, err := stmt.Exec("astaxie", "川林", "2012-12-09")
	chechErr(err)
	id, err := res.LastInsertId()
	chechErr(err)
	fmt.Println(id)
	//更新数据
	stmt,err=db.Prepare("update userinfo set username=? where uid=?")
	chechErr(err)
	stmt.Exec("astaxieupdata",id)
	chechErr(err)
	affect,err:=res.RowsAffected()//返回影响的数据条数
	chechErr(err)
	fmt.Println(affect,"affect")

	//查询数据
	rows,err:=db.Query("select * from userinfo")
	chechErr(err)
	for rows.Next(){
		var uid int
		var username string
		var department string
		var created string
		err=rows.Scan(&uid,&username,&department,&created)
		chechErr(err)
		fmt.Println(uid)
		fmt.Println(username)
		fmt.Println(department)
		fmt.Println(created)
	}
	//删除数据
	//stmt,err:=db.Prepare("delete from  userinfo where departname=?")
	//chechErr(err)
	//res,err:=stmt.Exec("jjj")
	//chechErr(err)
	//affect,err:=res.RowsAffected()
	//chechErr(err)
	//fmt.Println(affect)
	//db.Close()
}

func chechErr(err error)  {
	if err!=nil{
		fmt.Println(err)
		return
	}
}
/*
sql.Open()函数用来打开一个注册过的数据库驱动，go-sql-driver中注册了mysql这个数据库驱动，第二个参数是DSN(Data Source Name)，它是go-sql-driver定义的一些数据库链接和配置信息。它支持如下格式：

user@unix(/path/to/socket)/dbname?charset=utf8
user:password@tcp(localhost:5555)/dbname?charset=utf8
user:password@/dbname
user:password@tcp([de:ad:be:ef::ca:fe]:80)/dbname
db.Prepare()函数用来返回准备要执行的sql操作，然后返回准备完毕的执行状态。

db.Query()函数用来直接执行Sql返回Rows结果。

stmt.Exec()函数用来执行stmt准备好的SQL语句

我们可以看到我们传入的参数都是=?对应的数据，这样做的方式可以一定程度上防止SQL注入
 */