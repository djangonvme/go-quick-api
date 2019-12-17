package apis

import (
	"gin-api-common/services"
	"gin-api-common/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

//login api
func Login(c *gin.Context) {
	account := c.Query("mobile")
	pwd := c.Query("pwd")

	jwtToken, err := services.AppLogin(account, pwd)
	if err != nil {
		c.JSON(http.StatusOK, utils.FailResponse(err.Error()))
		return
	}
	c.JSON(http.StatusOK, utils.SuccessResponse(map[string]interface{}{"token": jwtToken}))
	return
}

//logout api
func Logout(c *gin.Context) {
	userId := c.GetInt64("login_user")
	err := services.AppLogout(userId)
	if err != nil {
		c.JSON(http.StatusOK, utils.FailResponse(err.Error()))
		return
	}
	c.JSON(http.StatusOK, utils.SuccessResponseSimple())
	return
}
