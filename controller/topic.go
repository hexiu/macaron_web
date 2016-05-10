package controller

import (
	"fmt"
	"gopkg.in/macaron.v1"
	"html/template"
	"log"
	"macaron/models"
)

func TopicHandler(ctx *macaron.Context, logger *log.Logger) {
	ctx.HTML(200, "topic_add")
}

func TopicAddHandler(ctx *macaron.Context, logger *log.Logger) {
	logg := ctx.Req.Method
	logger.Println(" \n\n\n ", string(logg), "\n\n\n\n\n\n")
	// ctx.Data["TEST"] = "ZYPC"

	// ctx.HTML(200, "topic_add")
	var Addtopic models.Topic
	Addtopic.Title = ctx.Req.Request.FormValue("title")
	Addtopic.Category = ctx.Req.Request.FormValue("category")
	Addtopic.Labels = ctx.Req.Request.FormValue("label")
	Addtopic.Condense = ctx.Req.Request.FormValue("condense")
	Addtopic.Content = ctx.Req.Request.FormValue("editor1")

	fmt.Println(Addtopic)
	models.AddTopic(Addtopic)
	ctx.Redirect("/topiclist", 301)
}

func TopicListHandler(ctx *macaron.Context, logger *log.Logger) (err error) {
	logg := ctx.Req.Method
	logger.Println(" \n\n\n ", string(logg), "\n\n\n\n\n\n")

	Topiclist, err := models.ListTopic()
	if err != nil {
		return err
	}

	ctx.Data["Topics"] = Topiclist
	ctx.Data["Hi"] = "hi hello"
	ctx.Data["Content"] = template.HTML(Topiclist[0].Content)

	ctx.HTML(200, "topic_list")
	return nil
}
