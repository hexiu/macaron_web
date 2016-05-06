package models

import (
	"github.com/go-xorm/xorm"
	// "github.com/mattn/go-sqlite3"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var DatabaseType string

var engine *xorm.Engine

type User struct {
	name string
	age  int
}

func RegisterDB() {
	var err error
	engine, err = xorm.NewEngine(DatabaseType, "root:axiu@127.0.0.1/mysql?charset=utf8")
	z, _ := engine.DBMetas()
	fmt.Println(z)
}
