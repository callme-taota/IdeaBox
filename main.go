package main

import (
	conf "idea-box/conf"
	"idea-box/database"
	"idea-box/server"
	"idea-box/tolog"
)

func main() {
	err := conf.InitConf()
	if err != nil {
		tolog.Log().Errorf("Config init %e", err).PrintAndWriteSafe()
		return
	}
	database.InitDB()
	server.InitServer()
}
