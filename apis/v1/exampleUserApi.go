package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/jangozw/gin-api-common/models"
	"github.com/jangozw/gin-api-common/params"
	"github.com/jangozw/gin-api-common/services"
	"github.com/jangozw/gin-api-common/utils"
)

func AddUser(c *gin.Context) {
	p := params.AddUser{}
	if err := c.ShouldBind(&p); err != nil {
		utils.Ctx(c).Fail(err)
		return
	}
	if err := models.AddUser(p.Name, p.Mobile, p.Pwd); err != nil {
		utils.Ctx(c).Fail(err)
		return
	}
	user, _ := models.FindUserByMobile(p.Mobile)
	utils.Ctx(c).Success(user)
	return
}

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
