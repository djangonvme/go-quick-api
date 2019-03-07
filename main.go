package main

import (
	"gin-api-common/config"
	"gin-api-common/database"
	"gin-api-common/routes"
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
