package v0

import (
	"github.com/jangozw/gin-api-common/params"
	"github.com/jangozw/gin-api-common/services"
	"github.com/jangozw/gin-api-common/utils"
)

//login api

func Login(c *utils.ApiContext) {
	p := params.Login{}
	if err := c.ShouldBind(&p); err != nil {
		c.Fail(err)
		return
	}
	jwtToken, err := services.AppLogin(p.Mobile, p.Pwd)
	if err != nil {
		c.Fail(err)
		return
	}
	c.Success(map[string]interface{}{"token": jwtToken})
	return
}

//logout api
func Logout(c *utils.ApiContext) {
	userId := c.GetLoginUid()
	if err := services.AppLogout(userId); err != nil {
		c.Fail(err)
		return
	}
	c.SuccessSimple()
	return
}
