package main

import (
	// "fmt"
	// "github.com/Unknwon/goconfig"
	"gopkg.in/macaron.v1"
	"macaron/modules/initConf"
)

func init() {
	initConf.InitConf()
}

func main() {
	m := macaron.Classic()
	m.Get("/", func() string {
		return "this is success!"
	})
	m.Run()
}
