package controller

import (
	"gopkg.in/macaron.v1"
	"log"
)

func TopicAddHandler(ctx *macaron.Context, logger *log.Logger) {
	logg := ctx.Req.Method
	logger.Println(" \n\n\n ", string(logg), "\n\n\n\n\n\n")
	// ctx.Data["TEST"] = "ZYPC"

	ctx.HTML(200, "topic_add")
	// return "the request is :" + ctx.Req.RequestURI
}
