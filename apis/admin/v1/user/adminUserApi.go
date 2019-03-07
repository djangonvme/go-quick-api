package adminUserApi

import (
	"github.com/gin-gonic/gin"
	"net/http"
	//"gin-api-common/models"
)
func List(c *gin.Context)  {
	//models.CreateToken()

	c.JSON(http.StatusOK, gin.H{"data":"adminuserlist"})

}
