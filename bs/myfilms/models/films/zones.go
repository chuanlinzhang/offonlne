package films

import (
	"gopkg.in/mgo.v2"
	"bs/myfilms/models/mongoDB"
)

/*
电影院放映厅
 */
type Zones struct {
	Zno       string `bson:"zno"`//放映厅编号
	Zname     string `bson:"zname"`//放映厅名字
	available string `bson:"available"`//该资料是否有效
	Latest    string `bson:"latest"`//最后一次更新时间
}
var collectionZones *mgo.Collection
func init() {
	dbname:=mongoDB.Dbname
	conn:=mongoDB.Dbsession.Copy()
	DB:=conn.DB(dbname)


	collectionZones=DB.C("zones")
}