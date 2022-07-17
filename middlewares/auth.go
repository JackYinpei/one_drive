package middlewares

import (
	"qqq_one_drive/controller"
	"qqq_one_drive/pkg/jwt"
	"strings"

	"github.com/gin-gonic/gin"
)

func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			controller.ResponseError(c, controller.CodeNeedLogin)
			c.Abort()
			return
		}

		part := strings.SplitN(authHeader, " ", 2)
		if !(len(part) == 2) && part[0] == "Bearer" {
			controller.ResponseError(c, controller.CodeInvalidToken)
			c.Abort()
			return
		}
		mc, err := jwt.ParseToken(part[1])
		if err != nil {
			controller.ResponseError(c, controller.CodeInvalidToken)
			c.Abort()
			return
		}

		c.Set(controller.CtxUserIDKey, mc.UserID)
		c.Next()
	}
}
