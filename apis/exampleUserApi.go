package apis

import (
	"gin-api-common/params"
	"gin-api-common/services"
	"gin-api-common/utils"
	"github.com/gin-gonic/gin"
	"net/http"
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
		c.JSON(http.StatusOK, utils.FailResponse(err.Error()))
		return
	}
	c.JSON(http.StatusOK, utils.SuccessResponse(data))
	return
}
