package api

import (
	"github.com/gin-gonic/gin"
	"idea-box/database"
	"idea-box/models"
	"idea-box/models/APIs/Request"
	"net/http"
)

func GetUserInfo(c *gin.Context) {
	var json Request.UserInfo
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, models.R(models.KErrorMissing, models.KReturnFalse, models.RDC{}))
		return
	}
	ok := models.ShouldCheckJSON(json, []string{"ID"})
	if ok != true {
		c.JSON(http.StatusOK, models.R(models.KErrorMissing, models.KReturnFalse, models.RDC{}))
		return
	}
	user, err := database.GetUserInfo(json.ID)
	if err != nil {
		c.JSON(http.StatusOK, models.R(models.KErrorMissing, models.KReturnFalse, models.RDC{}))
		return
	}
	c.JSON(http.StatusOK, models.Rs(models.KReturnMsgSuccess, models.KReturnTrue, user))
	return
}

func UserList(c *gin.Context) {
	userList, err := database.GetUserList()
	if err != nil {
		c.JSON(http.StatusOK, models.R(models.KErrorMissing, models.KReturnFalse, models.RDC{}))
		return
	}
	c.JSON(http.StatusOK, models.Rs(models.KReturnMsgSuccess, models.KReturnTrue, userList))
	return
}
