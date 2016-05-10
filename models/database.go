package models

import (
	"github.com/go-xorm/xorm"
	// "github.com/mattn/go-sqlite3"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	// "html/template"
	"time"
)

var DatabaseType string = "mysql"

var engine *xorm.Engine

var username string = "root"
var password string = "axiu"

type User struct {
	Id       int64
	Name     string
	Password string
	Flag     int64
	Created  time.Time `xorm:"index"`
}

type Topic struct {
	Id              int64
	Uid             int64
	Title           string
	Category        string
	Labels          string
	Condense        string `xorm:"SIZE(200)"`
	Content         string `xorm:"SIZE(5000)"`
	Attachment      string
	Created         time.Time `xorm:"index"`
	Updated         time.Time `xorm:"index"`
	Views           int64     `xorm:"index"`
	Author          string
	ReplyTime       time.Time `xorm:"index"`
	ReplyCount      int64
	ReplyLastUserId int64
}

type Reply struct {
	Id      int64
	Tid     int64
	Name    string
	Email   string
	Content string    `xorm:"SIZE(1000)"`
	Created time.Time `xorm:"index"`
}

func RegisterDB() error {
	// var err error
	engine, _ = xorm.NewEngine("mysql", username+":"+password+"@/test?charset=utf8")
	// if err != nil {
	// 	return err
	// }
	engine.Ping()
	if ok, _ := engine.IsTableExist("user"); !ok {
		// engine.CreateTables(new(User))
		engine.CreateTables(new(User))
	}
	if ok, _ := engine.IsTableExist("topic"); !ok {
		engine.CreateTables(new(Topic))
	}
	if ok, _ := engine.IsTableExist("reply"); !ok {
		engine.CreateTables(new(Reply))
	}
	defer engine.Close()
	return nil
}

func AddTopic(Addtopic Topic) (err error) {
	engine, _ = xorm.NewEngine("mysql", username+":"+password+"@/test?charset=utf8")

	// fmt.Println(Addtopic)
	Addtopic.Created = time.Now()
	engine.Ping()
	a, err := engine.Insert(Addtopic)
	fmt.Println(a, err)
	if err != nil {
		return err
	}
	return nil
}

func ListTopic() (listTopic []*Topic, err error) {
	listTopic = make([]*Topic, 0)
	topic := new(Topic)
	engine, _ = xorm.NewEngine("mysql", username+":"+password+"@/test?charset=utf8")
	engine.Get(topic)
	engine.Table("topic").OrderBy("Created").Find(&listTopic)

	fmt.Println(topic)
	fmt.Println(listTopic)

	return listTopic, nil
}
