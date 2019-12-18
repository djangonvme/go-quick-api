package apis

import (
	"gin-api-common/params"
	"gin-api-common/services"
	"gin-api-common/utils"
	"github.com/gin-gonic/gin"
)

func UserList(c *gin.Context) {
	mobile := c.Query("mobile")
	page := utils.StrToInt64(c.Query("page"))
	search := params.UserListSearch{
		Mobile: mobile,
		Page:   page,
	}
	data, err := services.GetUserList(search)
	if err != nil {
		utils.Response(c).Fail(err.Error())
		return
	}
	utils.Response(c).Success(data)
	return
}

func UserDetail(c *gin.Context) {
	utils.Response(c).SuccessSimple()
	return
}
