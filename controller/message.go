package controller

import (
	"gopkg.in/macaron.v1"
	"log"
)

func MessageHandler(ctx *macaron.Context, logger *log.Logger) string {
	logger.Println("message")
	return "the request is :" + ctx.Req.RequestURI
}
