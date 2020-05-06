package main

// 一些命令行工具
import (
	"fmt"
	"github.com/jangozw/gin-api-common/models"
)

func main() {
	//创建用户表一条记录
	err := models.AddUser("test", "1500000001", "123456")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("ok")
	}
}
