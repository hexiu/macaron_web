package main

import (
	// "fmt"
	// "github.com/Unknwon/goconfig"
	// "github.com/go-macaron/gzip"
	"gopkg.in/macaron.v1"
	// "log"
	"macaron/controller"
	"macaron/models"
	"macaron/modules/initConf"
)

func init() {
	models.RegisterDB()
	initConf.InitConf()
	// fmt.Println(conf)

}

func main() {
	m := macaron.Classic()
	m.Use(macaron.Renderer())
	m.Use(macaron.Logger())
	m.Use(macaron.Recovery())

	m.Get("/1", func() string {
		return "this is one."
	})
	m.Get("/", controller.HomeHandler)
	m.Get("/test", controller.TestHandler)
	m.Get("/message", controller.MessageHandler)
	m.Get("/topic", controller.TopicAddHandler)
	m.Run()
}
