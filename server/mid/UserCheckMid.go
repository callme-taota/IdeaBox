package mid

import (
	"github.com/gin-gonic/gin"
	"idea-box/conf"
	"idea-box/models"
	"net/http"
)

func UserCheckMid() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.URL.Path == "/api/auth" {
			c.Next()
			return
		}
		authKey, _ := c.Cookie("ideabox-auth-key")
		if !conf.CheckAuthKey(authKey) {
			c.JSON(http.StatusForbidden, models.R(models.KErrorPermissionDenied, models.KReturnFalse, models.RDC{}))
			c.Abort()
			return
		}
		userID := conf.GetAuthKeyID(authKey)
		c.Set("userID", userID)
		c.Next()
		return
	}
}
