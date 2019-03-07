package main

import (
	"gin-api-common/database"
	"gin-api-common/models"
)

func main()  {
	defer database.CloseDb()

	//创建所有定义好的表，都在models里
	models.CreateAllTable()

}









