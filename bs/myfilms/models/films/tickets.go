package films

import (
	"gopkg.in/mgo.v2"
	"bs/myfilms/models/mongoDB"

	"time"
	"strconv"
	"gopkg.in/mgo.v2/bson"


)

/*
影票信息
 */
type Tickets struct {
	Tno     int     `bson:"tno"`      //影票编号
	FilmsNo string  `bson:"films_no"` //影片编号
	Zno     int     `bson:"zno"`      //放映厅编号
	NameCh  string  `bson:"name_ch"`  //中文片名
	NameEn  string  `bson:"name_en"`  //英文片名
	Amount  float64 `bson:"amount"`   //订单金额
	StartTime  string `bson:"start_time"`  //放映时间
	FilmLength int    `bson:"film_length"` //片长
	Seat       string `bson:"seat"`        //座位号
	Available  bool   `bson:"availble"`    //资料是否有效
	Latest     string `bson:"latest"`      //最后一次更新时间
}

var collectionTickets *mgo.Collection

func init() {
	dbname := mongoDB.Dbname
	conn := mongoDB.Dbsession.Copy()
	DB := conn.DB(dbname)

	collectionTickets = DB.C("tickets")
}
func MakeTickets(films Films) bool {
	n := sumTickets()
	h := strconv.Itoa(14)
	for k := 1; k <= 2; k++ {
		for i := 1; i <= 5; i++ {
			for j := 1; j <= 6; j++ {
				tickets := &Tickets{
					Tno:        n+1,
					FilmsNo:films.FilmsNo,
					Zno:        k,
					NameEn:     films.NameEn,
					NameCh:     films.NameCh,
					Amount:     films.Pricefull,
					StartTime:  time.Now().Format("2006-01-02") + h + ":00:00",
					FilmLength: films.FilmLength,
					Seat:       strconv.Itoa(i) + "排" + strconv.Itoa(j) + "号",
					Available:  true,
					Latest:     time.Now().Format("2006-01-02 15:04:05"),
				}
				n++
				err := collectionTickets.Insert(tickets)
				if err != nil {
					return false
				}
			}

		}
		h1, _ := strconv.Atoi(h)
		h1++
		h = strconv.Itoa(h1)
	}

	return true
}
func DeleteTickets(FilmsNo int) bool {
	err := collectionTickets.Remove(bson.M{"tno": FilmsNo})
	if err != nil {
		return false
	}
	return true
}
func LookTickets1(FilmsNo string) []Tickets {
	list := make([]Tickets, 0) //定义一个切片
	collectionTickets.Find(bson.M{"films_no": FilmsNo,"zno":1}).All(&list)
	return list
}
func LookTickets2(FilmsNo string) []Tickets {
	list := make([]Tickets, 0) //定义一个切片
	collectionTickets.Find(bson.M{"films_no": FilmsNo,"zno":2}).All(&list)
	return list
}

func sumTickets() int {
	n, err := collectionTickets.Count()
	if err != nil {
		return -1
	}
	return n
}
