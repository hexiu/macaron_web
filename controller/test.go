package controller

import (
	"gopkg.in/macaron.v1"
	"log"
)

func TestHandler(ctx *macaron.Context, logger *log.Logger) {
	logger.Println("hello")

	ctx.Data["TEST"] = "ZYPC"
	ctx.Data["CESHI"] = "ZYPC"
	ctx.Data["NUM1"] = 300
	ctx.Data["NUM2"] = 200
	ctx.Data["NUM3"] = 500

	ctx.HTML(200, "test")
	// return "the request is :" + ctx.Req.RequestURI
}
