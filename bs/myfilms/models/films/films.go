package films

import (
	"gopkg.in/mgo.v2"

	"bs/myfilms/models/mongoDB"
	"time"
	"gopkg.in/mgo.v2/bson"
)

/*
影片信息
 */
type Films struct {
	FilmsNo    string  `bson:"films_no" form:"filmsNo"`    //影片编号
	NameCh     string  `bson:"name_ch" form:"nameCh"`     //中文片名
	NameEn     string  `bson:"name_en" form:"nameEn"`     //英文片名
	Director   string  `bson:"director" form:"director"`    //导演
	Actor      string  `bson:"actor" form:"actor"`       //主演
	Intro      string  `bson:"intro" form:"Intro"`       //影片简介
	FilmLength int     `bson:"film_length" form:"filmLength"` //片长
	Pricefull  float64 `bson:"pricefull" form:"Pricefull"`   //全票价
	Pricest    float64 `bson:"pricest" form:"Pricest"`     //学生票价
	Hot        int     `bson:"hot" form:"hot"`        //热度
	New        int     `bson:"new" form:"new"`         //新片
	Available  bool    `bson:"available"`   //该影片资料是否有效
	Latest     string  `bson:"latest"`      //最后一次更新时间
}

var collectionFilms *mgo.Collection

func init() {
	dbname := mongoDB.Dbname
	conn := mongoDB.Dbsession.Copy()
	DB := conn.DB(dbname)

	collectionFilms = DB.C("films")
}
//func Addfilms(FilmsNo, NameCh, NameEn, Director, Actor, Intro string, FilmLength, Hot, New int, Pricefull, Pricest float64) bool {
//	films := &Films{
//		FilmsNo:    FilmsNo,
//		NameCh:     NameCh,
//		NameEn:     NameEn,
//		Director:   Director,
//		Actor:      Actor,
//		Intro:      Intro,
//		FilmLength: FilmLength,
//		Pricefull:  Pricefull,
//		Pricest:    Pricest,
//		Hot:Hot,
//		New:New,
//		Available:  true,
//		Latest:     time.Now().Format("2006-01-02 15:04:05"),
//	}
//	err := collectionFilms.Insert(films)
//	if err != nil {
//		return false
//	}
//	return true
//}
func Addfilms1(films Films) bool {
	films.Available=true
	films.Latest=time.Now().Format("2006-01-02 15:04:05")
	err := collectionFilms.Insert(films)
	if err != nil {
		return false
	}
	return true
}
func FilmsTreeGrid() []Films {
	//查询所有
	list := make([]Films, 0) //定义一个切片
	collectionFilms.Find(nil).All(&list)
	return list
}
func Changefilms(FilmsNo string) *Films {
	films := &Films{}
	err := collectionFilms.Find(bson.M{"films_no": FilmsNo}).One(films)
	if err != nil {
		return nil
	}
	return films
}
func Deletefilms(FilmsNo string) bool {
	err := collectionFilms.Remove(bson.M{"films_no": FilmsNo})
	if err != nil {
		return false
	}
	return true
}


func Sumfilms() int {
	n, err := collectionFilms.Count()
	if err != nil {

		return -1
	}
	return n
}
func Hotfilms() []Films {
	list := make([]Films, 0) //定义一个切片
	err:=collectionFilms.Find(nil).Sort("-hot").All(&list)
	if err!=nil {
		return nil
	}
	return list
}
func Newfilms() []Films {
	list := make([]Films, 0) //定义一个切片
	err:=collectionFilms.Find(nil).Sort("-new").All(&list)
	if err!=nil{
		return nil
	}
	return list
}