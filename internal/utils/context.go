package utils

import (
	"context"
	//"github.com/closetotheworld/gqlgen-gin-practice/pkg/jwt"
	"github.com/gin-gonic/gin"
	//"net/http"
)

type contextKey string

func (c contextKey) String() string {
	return string(c)
}

const (
	userCtxKey = contextKey("UserContextKey")
	ginCtxKey  = contextKey("GinContextKey")
)

func GinContextToContextMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.WithValue(c.Request.Context(), ginCtxKey, c)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

//func JwtMiddleware() gin.HandlerFunc {
//	return func(c *gin.Context) {
//		jwtStr := c.GetHeader("Authorization")
//
//		if jwtStr == "" {
//			c.Next()
//			return
//		}
//
//		uid, err := jwt.ParseToken(jwtStr)
//		if err != nil {
//			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"errors": []gin.H{
//				{"message": "invalid token"},
//			}})
//			return
//		}
//
//	}
//}
