package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"oauth2/httpType"
)

// 定义中间件函数，用于检查 Bearer Token
func AuthMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, httpType.Response{
				Code: 200011,
				Data: nil,
				Msg:  "Missing Authorization header",
			})
			c.Abort()
			return
		}

		// 假设 Token 格式为 "Bearer <token>"
		// 实际应用中你需要从字符串中提取真正的 Token
		token := authHeader[len("Bearer "):]
		c.Set("access_token", token)

		c.Next()
	}
}

func ErrorHandlingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		// 检查处理函数是否返回了错误
		if len(c.Errors) > 0 {
			// 可以在这里记录日志、返回自定义错误响应等
			for _, err := range c.Errors {
				fmt.Println("Error:", err.Error())
			}
			c.AbortWithStatusJSON(http.StatusInternalServerError, httpType.Response{
				Code: 200001,
				Data: nil,
				Msg:  c.Errors[0].Error(), //return the last error
			})
		}
	}
}
