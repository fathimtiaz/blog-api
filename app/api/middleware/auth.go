package middleware

import (
	"blog-api/internal/domain"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func SetAuthdUserCtx(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		var tokenStr string

		authHeader := strings.Split(c.GetHeader("Authorization"), "Bearer ")
		if len(authHeader) > 1 {
			tokenStr = authHeader[1]
		} else {
			c.Next()
			return
		}

		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(secret), nil
		})
		if err != nil {
			log.Println(err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, err)
			return
		}
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			if email, err := claims.GetSubject(); err == nil {
				c.Set(domain.AuthdUserEmailCtx, email)
				c.Next()
			}

			return
		} else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, "invalid token")
			return
		}
	}
}
