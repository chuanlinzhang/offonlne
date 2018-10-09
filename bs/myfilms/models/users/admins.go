package users

/*
管理员
 */
import (
	_ "gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"
	"github.com/astaxie/beego/validation"
	"gopkg.in/mgo.v2/bson"

	"time"
	"bs/myfilms/models/mongoDB"
	"github.com/astaxie/beego/logs"
	"fmt"
	"log"
)

type Admins struct {
	AdminLogin    string `bson:"admin_login" form:"admin_login"`    //管理员登陆账号
	AdminPwd      string `bson:"admin_pwd" form:"admin_pwd1"`      //管理员登陆密码
	AdminNickName string `bson:"admin_nickname" form:"admin_nickname"` //管理员昵称
	AdminName     string `bson:"admin_name" form:"admin_name"`     //管理员实名
	AdminMobile   string `bson:"admin_mobile" form:"admin_mobile" valid:"Mobile"`   //管理员手机号码
	AdminSex      string `bson:"admin_sex" form:"admin_sex"`      //管理员性别
	AdminEmail    string `bson:"admin_email" form:"admin_email"`    //管理员Email
	Available     bool   `bson:"available"`                         //资料是否有效
	Latest        string `bson:"latest"`                           //最后一次更新时间
}

var collectionAdmins *mgo.Collection

func init() {
	dbname := mongoDB.Dbname
	conn := mongoDB.Dbsession.Copy()
	DB := conn.DB(dbname)
	collectionAdmins = DB.C("admins")
}
func RegisterA1(admins1 Admins) bool {
	admins := &Admins{}
	err := collectionAdmins.Find(bson.M{"admin_login": admins1.AdminLogin}).One(admins)
	if err == nil {
		logs.Info("管理员已存在")
		return false
	}
	admins1.Available=true
	admins1.Latest=time.Now().Format("2006-01-02 15:04:05")
	//验证用户输入信息格式是否正确
	valid := validation.Validation{}
	b, err := valid.Valid(admins1)
	if err != nil&&!b {
		log.Println("邮箱或者手机号码格式不符合")
		return false
	}
	err = collectionAdmins.Insert(admins1)

	if err != nil {
		return false
	}
	return true
}
//func RegisterA(adminLogin, adminPwd, adminNickname, adminName, adminMobile, adminSex, adminEmail string) bool {
//	admins := &Admins{}
//	err := collectionAdmins.Find(bson.M{"admin_login": adminLogin}).One(admins)
//	if err == nil {
//		logs.Info("管理员已存在")
//		return false
//	}
//
//	admins = &Admins{
//		AdminLogin:    adminLogin,
//		AdminPwd:      adminPwd,
//		AdminNickName: adminNickname,
//		AdminName:     adminName,
//		AdminMobile:   adminMobile,
//		AdminSex:      adminSex,
//		AdminEmail:    adminEmail,
//		Available:     true,
//		Latest:        time.Now().Format("2006-01-02 15:04:05"),
//	}
//	//验证用户输入信息格式是否正确
//	valid := validation.Validation{}
//	_, err = valid.Valid(admins)
//	if err != nil {
//		return false
//	}
//	err = collectionAdmins.Insert(admins)
//
//	if err != nil {
//		return false
//	}
//	return true
//}
func LoginA(adminLogin string) *Admins {
	admins := &Admins{}
	err := collectionAdmins.Find(bson.M{"admin_login": adminLogin}).One(admins)
	if err != nil {
		return nil
	}
	return admins
}
func ChangeA(adminLogin, adminNickname, adminName, adminMobile, adminSex, adminEmail string) bool {
	err := collectionAdmins.Update(bson.M{"admin_login": adminLogin}, bson.M{"$set": bson.M{
		"admin_nickname": adminNickname,
		"admin_name":     adminName,
		"admin_mobile":   adminMobile,
		"admin_sex":      adminSex,
		"admin_email":    adminEmail,
	}})
	if err != nil {
		fmt.Println("err", err)
		return false
	}
	return true
}
func ChangePwdA(adminLogin, pwd string) bool {
	err := collectionAdmins.Update(bson.M{"admin_login": adminLogin}, bson.M{"$set": bson.M{"admin_pwd": pwd}})
	if err != nil {
		return false
	}
	return true
}
