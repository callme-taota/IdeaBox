package mid

import (
	"github.com/gin-gonic/gin"
	"idea-box/tolog"
)

func LogMid(c *gin.Context) {
	tolog.Log().Infof("Gin RequestMethod: %s, URL Path: %s", c.Request.Method, c.Request.URL.Path).PrintAndWriteSafe()
	c.Next()
}
