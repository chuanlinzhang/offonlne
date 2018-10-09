package films

import (
	"gopkg.in/mgo.v2"
	"bs/myfilms/models/mongoDB"
)

/*
电影院放映厅座位表
 */
type Seats struct {
	Zno       string `bson:"zno"`//放映厅编号
	Snox      string `bson:"snox"`//座位的排
	Snoy      string `bson:"snoy"`//座位的列
	Rank      string `bson:"rank"`//座位等级
	Available string `bson:"available"`//该资料是否有效
	Latest    string `bson:"latest"`//最后一次更新时间
}
var collectionSeats *mgo.Collection
func init() {
	dbname:=mongoDB.Dbname
	conn:=mongoDB.Dbsession.Copy()
	DB:=conn.DB(dbname)

	collectionSeats=DB.C("seats")
}