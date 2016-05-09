package controller

import (
	"gopkg.in/macaron.v1"
	"log"
)

func HomeHandler(ctx *macaron.Context, logger *log.Logger) {
	logg := ctx.Req.Method
	logger.Println(" \n\n\n ", string(logg), "\n\n\n\n\n\n")
	// ctx.Data["TEST"] = "ZYPC"
	ctx.Data["NUM1"] = 300
	ctx.Data["NUM2"] = 200
	ctx.Data["NUM3"] = 500

	ctx.HTML(200, "index")
}
