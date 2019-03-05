package userApi

import (
	"github.com/gin-gonic/gin"
	"github.com/jangozw/gintest/models"
	"strconv"
	"net/http"
	"github.com/jangozw/gintest/apis"
	"time"
)

//接口返回的用户数据结构
type user struct {
	Id uint
	Name string
	Birthday time.Time
}

//用户列表接口






func List(c *gin.Context) {
	page,_ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize,_ := strconv.Atoi(c.DefaultQuery("pagesize", "3"))
	userModel := models.User{}
	lists := userModel.List(page, pageSize)
	users := getUsers(lists)
	//return OutputSuccess(c, users)
	c.JSON(http.StatusOK, apis.SuccessFormat(users))
	return
}

//用户详情接口
func Detail(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		c.JSON(http.StatusOK, apis.ErrorFormat(apis.CODE_ERROR, err))
		//不往下执行怎么做？
		return
	}
	if id <= 0 {
		c.JSON(http.StatusOK, apis.ErrorFormat(apis.CODE_ERROR, "id 不合法"))
		//不往下执行怎么做？
		return
	}
	userModel := models.User{}
	user := userModel.Detail(id)
	detail := getDetailItem(user)
	c.JSON(http.StatusOK, apis.SuccessFormat(detail))
	return
}

//从Model筛选出如何上述格式的列表数据，注意这种二维切片的实现方法
func getUsers(lists []models.User) []user  {
	users := make([]user, 0)
	for _,v:= range lists {
		u := getDetailItem(v)
		users = append(users, u)
	}
	return users
}
//detail
func getDetailItem(m models.User)  user{
	return user{
		m.ID,
		m.Name,
		m.Birthday,
	}
}