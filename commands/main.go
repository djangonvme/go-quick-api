package main

import (
	"github.com/jangozw/gintest/database"
	"github.com/jangozw/gintest/models"
)

func main()  {
	defer database.CloseDb()

	models.CreateAllTable()

}






