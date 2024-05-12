package server

import (
	"idea-box/conf"
	"idea-box/server/api"
)

func LinkUser() {
	userGroup := Server.Group("/user")

	userGroup.GET("/", api.GetUserInfo)
	userGroup.GET("/users", api.UserList)
}

func LinkIdea() {
	ideaGroup := Server.Group("/idea")

	ideaGroup.GET("/ideas", api.GetIdeaList)
	ideaGroup.POST("/", api.CreateIdea)
	ideaGroup.POST("/update", api.UpdateIdea)
	ideaGroup.POST("/name", api.UpdateIdeaName)
	ideaGroup.POST("/desc", api.UpdateIdeaDesc)
	ideaGroup.DELETE("/", api.DeleteIdea)
}

func LinkComment() {
	commentGroup := Server.Group("/comment")

	commentGroup.POST("/", api.CreateComment)
	commentGroup.DELETE("/", api.DeleteComment)
	commentGroup.GET("/", api.GetCommentList)
}

func LinkDebug() {
	if conf.Server.Model == "debug" {
		Server.GET("/"+conf.RandomKey, api.SetDebugKeyInCookie)
	}
	Server.POST("/auth", api.SetAuthCookie)
}
