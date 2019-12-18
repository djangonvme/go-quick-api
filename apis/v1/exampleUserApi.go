package v1

import (
	"gin-api-common/params"
	"gin-api-common/services"
	"gin-api-common/utils"
	"github.com/gin-gonic/gin"
)

func UserList(c *gin.Context) {
	//校验请求参数, 校验规则定义在params.SearchUserList{}的tag里
	search := params.SearchUserList{}
	if err := c.ShouldBind(&search); err != nil {
		utils.Ctx(c).Fail(err)
		return
	}
	//校验参数成功后自动赋值给结构体
	if data, err := services.GetUserList(search); err != nil {
		utils.Ctx(c).Fail(err)
		return
	} else {
		utils.Ctx(c).Success(data)
		return
	}
}

func UserDetail(c *gin.Context) {
	utils.Ctx(c).SuccessSimple()
	return
}
