package main

import (
	"github.com/jangozw/gintest/routes"
	"github.com/gin-gonic/gin"
	"github.com/jangozw/gintest/database"
)

func main() {
	defer database.CloseDb()
	r := gin.New()
	routes.InitApiRouter(r)
	routes.InitAdminRouter(r)
	routes.InitNoTokenRouter(r)
	r.Run(":8080")
}
