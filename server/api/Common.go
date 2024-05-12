package api

import (
	"github.com/gin-gonic/gin"
	"idea-box/conf"
	"idea-box/models"
	"idea-box/models/APIs/Request"
	"net/http"
)

func SetDebugKeyInCookie(c *gin.Context) {
	key := conf.RandomKey
	c.SetCookie("ideabox-debug-key", key, 3600*24*30, "/", "", false, true)
	c.JSON(http.StatusOK, models.R(models.KReturnMsgSuccess, models.KReturnTrue, models.RDC{}))
	return
}

func SetAuthCookie(c *gin.Context) {
	var json Request.AuthJSON
	if err := c.ShouldBind(&json); err != nil {
		c.JSON(http.StatusBadRequest, models.R(models.KReturnMsgError, models.KReturnFalse, models.RDC{}))
		return
	}
	ok := models.ShouldCheckJSON(json, []string{"Key"})
	if ok != true {
		c.JSON(http.StatusOK, models.R(models.KErrorMissing, models.KReturnFalse, models.RDC{}))
		return
	}
	if conf.CheckAuthKey(json.Key) {
		c.SetCookie("ideabox-auth-key", json.Key, 3600*24*30, "/", "", false, true)
		c.JSON(http.StatusOK, models.R(models.KReturnMsgSuccess, models.KReturnTrue, models.RDC{}))
		return
	}
	c.JSON(http.StatusOK, models.R(models.KErrorMissing, models.KReturnFalse, models.RDC{}))
	return
}
