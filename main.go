package main

import (
	// "fmt"
	// "github.com/Unknwon/goconfig"
	// "github.com/go-macaron/gzip"
	"gopkg.in/macaron.v1"
	"log"
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
	m.Get("/", myhandler)
	m.Get("/message", controller.MessageHandler)
	m.Run()

}

func myhandler(ctx *macaron.Context, logger *log.Logger) {
	logger.Println("hello\nhello\nhello\nhello\nhello\nhello\nhello\nhello\nhello\nhello\n")
	ctx.Data["TEST"] = "ZYPC"
	ctx.HTML(200, "index")
	// return "the request is :" + ctx.Req.RequestURI
}

func testhandler(ctx *macaron.Context, logger *log.Logger) string {
	logger.Println("hello")
	return "the request is :" + ctx.Req.RequestURI
}
