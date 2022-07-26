package middlewares

import (
	"qqq_one_drive/controller"
	"qqq_one_drive/pkg/jwt"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			controller.ResponseError(c, controller.CodeNeedLogin)
			c.Abort()
			return
		}

		// part := strings.SplitN(authHeader, " ", 2)
		// if !(len(part) == 2) && part[0] == "Bearer" {
		// 	zap.L().Error("前端传的Token 有点问题, 这里是拆分的问题")
		// 	controller.ResponseError(c, controller.CodeInvalidToken)
		// 	c.Abort()
		// 	return
		// }
		mc, err := jwt.ParseToken(authHeader)
		if err != nil {
			zap.L().Error("前端传的Token 有点问题，解析的时候的问题了就")
			controller.ResponseError(c, controller.CodeInvalidToken)
			c.Abort()
			return
		}

		c.Set(controller.CtxUserIDKey, mc.UserID)
		c.Next()
	}
}
