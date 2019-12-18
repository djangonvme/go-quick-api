package apis

import (
	"gin-api-common/services"
	"gin-api-common/utils"
	"github.com/gin-gonic/gin"
)

//login api
func Login(c *gin.Context) {
	account := c.Query("mobile")
	pwd := c.Query("pwd")
	jwtToken, err := services.AppLogin(account, pwd)
	if err != nil {
		utils.Response(c).Fail(err)
		return
	}
	utils.Response(c).Success(map[string]interface{}{"token": jwtToken})
	return
}

//logout api
func Logout(c *gin.Context) {
	userId := getLoginUid(c)
	err := services.AppLogout(userId)
	if err != nil {
		utils.Response(c).Fail(err)
		return
	}
	utils.Response(c).SuccessSimple()
	return
}
