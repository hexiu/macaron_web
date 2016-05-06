package main

import (
	// "fmt"
	// "github.com/Unknwon/goconfig"
	// "github.com/go-macaron/gzip"
	"gopkg.in/macaron.v1"
	// "log"
	"macaron/controller"
	"macaron/modules/initConf"
)

func init() {
	initConf.InitConf()
}

func main() {
	m := macaron.Classic()
	m.Use(macaron.Renderer())
	m.Use(macaron.Logger())
	m.Use(macaron.Recovery())

	// m.Use(gzip.Gziper())
	// log.SetOutput(w)
	// m.Get("/", func() string {
	// 	return "this is /"
	// })
	m.Get("/", controller.HomeHandler)
	m.Get("/test", controller.TestHandler)
	m.Get("/message", controller.MessageHandler)
	m.Get("/edit", controller.EditHandler)
	// m.Get("/edit", )
	m.Run()
}
