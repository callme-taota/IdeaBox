package api

import (
	"github.com/gin-gonic/gin"
	"idea-box/conf"
	"idea-box/models"
	"net/http"
)

func SetDebugKeyInCookie(c *gin.Context) {
	key := conf.RandomKey
	c.SetCookie("ideabox-debug-key", key, 3600*24*30, "/", "", false, true)
	c.JSON(http.StatusOK, models.R(models.KReturnMsgSuccess, models.KReturnTrue, models.RDC{}))
	return
}
