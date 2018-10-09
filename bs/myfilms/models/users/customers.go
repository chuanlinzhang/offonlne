package users

/*
会员
 */
import (
	_ "gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"

	"time"
	"github.com/astaxie/beego/validation"
	"gopkg.in/mgo.v2/bson"
	"bs/myfilms/models/mongoDB"
	"github.com/astaxie/beego/logs"
	"bs/myfilms/models/sendemail"
)

type Customers struct {
	LoginName string  `bson:"login_name" form:"cname"` //登陆账号
	Pwd       string  `bson:"pwd" form:"cpwd1"`        //登录密码
	Email     string  `bson:"email" form:"cemail" valid:"Email; MaxSize(100)"`      //Email
	Nickname  string  `bson:"nick_name" form:"nick_name"`  //昵称
	Name      string  `bson:"name" form:"name"`       //会员实名
	Mobile    string  `bson:"mobile" form:"mobile" valid:"Mobile"`     //手机号码
	Sex       string  `bson:"sex" form:"sex"`        //性别
	Spent     float64 `bson:spent`                     //总消费金额
	Balance   float64 `bson:"balance"`                 //余额
	Memlv     int32   `bson:"menlv"`                   //会员等级
	Available bool    `bson:"available"`               //资料是否有效
	Latest    string  `bson:"latest"`                  //最后一次更新时间
}

var collectionCustomers *mgo.Collection

func init() {
	dbname := mongoDB.Dbname
	conn := mongoDB.Dbsession.Copy()
	DB := conn.DB(dbname)
	collectionCustomers = DB.C("customers")
}
func Register1(customers1 Customers) bool {
	customers := &Customers{}
	err := collectionCustomers.Find(bson.M{"login_name": customers1.LoginName}).One(customers)
	if err == nil {
		logs.Info("用户已存在")
		return false
	}
	customers1.Spent=0
	customers1.Balance=0
	customers1.Memlv=1
	customers1.Available=true
	customers1.Latest=time.Now().Format("2006-01-02 15:04:05")
	//验证用户输入信息格式是否正确
	valid := validation.Validation{}
	_, err = valid.Valid(customers1)
	if err != nil {
		return false
	}
	err = collectionCustomers.Insert(customers1)
	if err != nil {
		return false
	}
	return true
}
//func Register(loginName, pwd, email, nickName, name, moile, sex string) bool {
//
//	customers := &Customers{}
//	err := collectionCustomers.Find(bson.M{"login_name": loginName}).One(customers)
//	if err == nil {
//		logs.Info("用户已存在")
//		return false
//	}
//	customers = &Customers{
//		LoginName: loginName,
//		Pwd:       pwd,
//		Email:     email,
//		Nickname:  nickName,
//		Name:      name,
//		Mobile:    moile,
//		Sex:       sex,
//		Spent:     0,
//		Balance:   0,
//		Memlv:     1,
//		Available: true,
//		Latest:    time.Now().Format("2006-01-02 15:04:05"),
//	}
//	//验证用户输入信息格式是否正确
//	valid := validation.Validation{}
//	_, err = valid.Valid(customers)
//	if err != nil {
//		return false
//	}
//	err = collectionCustomers.Insert(customers)
//	if err != nil {
//		return false
//	}
//	return true
//}
func Login(loginName string) *Customers {
	customers := &Customers{}
	err := collectionCustomers.Find(bson.M{"login_name": loginName}).One(customers)
	if err != nil {
		return nil
	}
	return customers
}
func Change(loginName, email, nickName, name, moile, sex string) bool {
	err := collectionCustomers.Update(bson.M{"login_name": loginName}, bson.M{"$set": bson.M{
		"email":     email,
		"nick_name": nickName,
		"name":      name,
		"mobile":    moile,
		"sex":       sex,
	}})
	if err != nil {
		return false
	}
	go sendemail.ChangeInfoSE(loginName,email)
	return true
}
func ChangePwd(loginName,email, newPwd1 string) bool {
	err := collectionCustomers.Update(bson.M{"login_name": loginName}, bson.M{"$set": bson.M{"pwd": newPwd1}})
	if err != nil {
		return false
	}

	go sendemail.ChangePwdSE(loginName,email)
	return true
}
func TopUp(loginName, pwd string, balance float64) bool {
	err := collectionCustomers.Update(bson.M{"login_name": loginName, "pwd": pwd}, bson.M{"$set": bson.M{"balance": balance}})
	if err != nil {
		return false
	}
	return true
}

//获取所以的用户
func CustomersTreeGrid() []Customers {
	//查询所有
	list := make([]Customers, 0) //定义一个切片
	collectionCustomers.Find(nil).All(&list)
	return list
}

func DelCus(loginName string) bool {
	err := collectionCustomers.Remove(bson.M{"login_name": loginName})
	if err != nil {
		return false
	}
	return true
}

////将资源列表转成treegrid格式
//func resourceList2TreeGrid(list []*Customers) []*Customers {
//	result := make([]*Customers, 0)
//	for _, item := range list {
//		if item.Parent == nil || item.Parent.Id == 0 {
//			item.Level = 0
//			result = append(result, item)
//			result = resourceAddSons(item, list, result)
//		}
//	}
//	return result
//}
