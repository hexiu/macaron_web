package controller

import (
	"gopkg.in/macaron.v1"
	"log"
)

func TestHandler(ctx *macaron.Context, logger *log.Logger) {
	logger.Println("hello")

	ctx.Data["TEST"] = "ZYPC"
	ctx.HTML(200, "test")
}
