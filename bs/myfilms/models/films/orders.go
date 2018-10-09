package films

import (
	"gopkg.in/mgo.v2"
	"bs/myfilms/models/mongoDB"
	//"bs/myfilms/models/users"
	"time"
	"gopkg.in/mgo.v2/bson"
	"bs/myfilms/models/users"
	"fmt"
)

/*
订单信息
 */
type Orders struct {
	OrderNo       int     `bson:"order_no"`       //订单编号
	OrderDate     string  `bson:"order_date"`     //下单时间
	OrderIdentify string  `bson:"order_identify"` //取票码
	LoginName     string  `bson:"login_name"`  //会员账号
	Name          string  `bson:"name"`       //会员实名
	Amount        float64 `bson:"amount"`     //订单金额
	Mobile        string  `bson:"mobile"`     //付款方式
	Status        string  `bson:"status"`     //订单状态
	NameCh        string  `bson:"name_ch"`    //中文片名
	StartTime     string  `bson:"start_time"` //放映时间
	Seat          string  `bson:"seat"`       //座位号
	Available     bool    `bson:"availble"`   //资料是否有效
	Latest        string  `bson:"latest"`     //最后一次更新时间
}

var collectionOrders *mgo.Collection

func init() {
	dbname := mongoDB.Dbname
	conn := mongoDB.Dbsession.Copy()
	DB := conn.DB(dbname)

	collectionOrders = DB.C("orders")
}
func AddOrders(customers *users.Customers, NameCh, StartTime, Seat string, Amount1 float64) bool {
	n := sumOrders() + 1
	orders := &Orders{
		OrderNo:       n,
		OrderDate:     time.Now().Format("2006-01-02 15:04:05"),
		OrderIdentify: "123456",
		LoginName:customers.LoginName,
		Name:          customers.Name,
		Amount:        Amount1,
		Mobile:        "微信付款",
		Status:        "未观看",
		NameCh:        NameCh,
		StartTime:     StartTime,
		Seat:          Seat,
		Available:     true,
		Latest:        time.Now().Format("2006-01-02 15:04:05"),
	}
	err := collectionOrders.Insert(orders)
	if err != nil {
		return false
	}
	return true
}
func sumOrders() int {
	n, _ := collectionOrders.Count()
	return n
}
func LookOrders(LoginName string) []Orders {
	list := make([]Orders, 0)
	err:=collectionOrders.Find(bson.M{"login_name":LoginName}).All(&list)
	if err!=nil{
		return nil
	}
	return list
}
func DelOrders(OrderNo int) bool {
	err:=collectionOrders.Remove(bson.M{"order_no":OrderNo})

	if err!=nil{
		fmt.Println(err,OrderNo)
		return false
	}
	return true
}