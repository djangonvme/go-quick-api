package v0

import (
	"fmt"
	"github.com/jangozw/gin-api-common/models"
	"github.com/jangozw/gin-api-common/params"
	"github.com/jangozw/gin-api-common/services"
	"github.com/jangozw/gin-api-common/utils"
	"time"
)

func AddUser(c *utils.ApiContext) {
	p := params.AddUser{}
	if err := c.ShouldBind(&p); err != nil {
		c.Fail(err)
		return
	}
	if err := models.AddUser(p.Name, p.Mobile, p.Pwd); err != nil {
		c.Fail(err)
		return
	}
	user, _ := models.FindUserByMobile(p.Mobile)
	c.Success(user)
	return
}

func UserList(c *utils.ApiContext) {
	//校验请求参数, 校验规则定义在params.SearchUserList{}的tag里
	search := params.SearchUserList{}
	if err := c.ShouldBind(&search); err != nil {
		c.Fail(err)
		return
	}
	//校验参数成功后自动赋值给结构体
	if data, err := services.GetUserList(search); err != nil {
		c.Fail(err)
		return
	} else {
		c.Success(data)
		return
	}
}

func UserDetail(c *utils.ApiContext) {
	c.SuccessSimple()
	return
}

func TimeOutOperation(c *utils.ApiContext) {
	fmt.Println("超时5秒，表演开始...")
	time.Sleep(4 * time.Second)
	select {
		case <-c.Request.Context().Done():
			// 超时会在中间件自动输出，这里不用输出
			return
		default:
			c.SuccessSimple()
	}
	fmt.Println("end'")
}
