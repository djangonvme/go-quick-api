package main

import (
	"github.com/jangozw/gintest/routes"
	"github.com/jangozw/gintest/database"
	"github.com/jangozw/gintest/config"
	"github.com/gin-gonic/gin"
)


func main() {
	defer database.CloseDb()
	r := gin.New()
	routes.InitApiRouter(r)
	routes.InitAdminRouter(r)
	routes.InitNoTokenRouter(r)
	listen := config.GetValue("env", "listen")
	r.Run(":"+listen)
}
