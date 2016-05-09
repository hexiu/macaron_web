package models

import (
	"github.com/go-xorm/xorm"
	// "github.com/mattn/go-sqlite3"
	// "fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

var DatabaseType string = "mysql"

var engine *xorm.Engine

type User struct {
	Id       int64
	Name     string
	Password string
	Created  time.Time `xorm:"index"`
}

type Topic struct {
	Id              int64
	Uid             int64
	Title           string
	Category        string
	Labels          string
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
	engine, _ = xorm.NewEngine("mysql", "root:axiu@/test?charset=utf8")
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
