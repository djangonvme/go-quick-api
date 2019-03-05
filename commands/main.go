package main

import (
	"github.com/jangozw/gintest/database"
	"github.com/jangozw/gintest/models"
)

func main()  {
	defer database.CloseDb()

	//创建所有定义好的表，都在models里
	models.CreateAllTable()

}









