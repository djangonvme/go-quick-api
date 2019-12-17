package main

import (
	"fmt"
	"gin-api-common/configs"
	"gin-api-common/routes"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	r := gin.New()
	routes.InitApiRouters(r)
	listen, err := configs.Get("server", "listen")
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("gin try running on http://127.0.0.1:" + listen)

	if err := r.Run(":" + listen); err != nil {
		log.Println("gin running error:", err.Error())
	}

}
