package middlewares

import (
	gincontext "FeasOJ/internal/gin"
	"FeasOJ/internal/utils/sql"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
)

func PermissionChecker() gin.HandlerFunc {
	return func(c *gin.Context) {
		encodedUsername := c.GetHeader("Username")
		username, _ := url.QueryUnescape(encodedUsername)
		if sql.SelectUserInfo(username).Role != 1 {
			c.JSON(http.StatusForbidden, gin.H{"message": gincontext.GetMessage(c, "forbidden")})
			c.Abort()
			return
		}
		c.Next()
	}
}
