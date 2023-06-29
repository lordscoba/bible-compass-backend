package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lordscoba/bible_compass_backend/internal/config"
	"github.com/lordscoba/bible_compass_backend/utility"
)

func AuthMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {
		clientToken := c.Request.Header.Get("token")

		if clientToken == "" {
			rd := utility.BuildErrorResponse(http.StatusInternalServerError, "error", "No authorization header was provided", "No authorization", nil)
			c.JSON(http.StatusInternalServerError, rd)
			c.Abort()
			return
		}
		secretkey := config.GetConfig().Server.Secret

		claims, err := utility.ValidateToken(secretkey, clientToken)

		if err != "" {
			rd := utility.BuildErrorResponse(http.StatusInternalServerError, "error", err, err, nil)
			c.JSON(http.StatusInternalServerError, rd)
			c.Abort()
			return
		}
		c.Set("email", claims.Email)
		c.Set("userame", claims.Username)
		c.Set("uid", claims.Uid)
		c.Set("type", claims.Type)

		// Continue to the next middleware or handler
		c.Next()
	}

}
