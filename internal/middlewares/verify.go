package middlewares

import (
	gincontext "FeasOJ/internal/gin"
	"FeasOJ/internal/utils"
	"FeasOJ/internal/utils/sql"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
)

func HeaderVerify() gin.HandlerFunc {
	return func(c *gin.Context) {
		var User string
		encodedUsername := c.GetHeader("Username")
		username, err := url.QueryUnescape(encodedUsername)
		token := c.GetHeader("Authorization")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": gincontext.GetMessage(c, "userNotFound")})
			c.Abort()
			return
		}
		if utils.IsEmail(username) {
			User = sql.SelectUserByEmail(username).Username
		} else {
			User = username
		}
		if !utils.VerifyToken(User, token) {
			c.JSON(http.StatusUnauthorized, gin.H{"message": gincontext.GetMessage(c, "unauthorized")})
			c.Abort()
			return
		}
		c.Next()
	}
}
