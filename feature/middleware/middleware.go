package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func AuthMiddleware(secretKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		//Take header name authorize
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Need to login as admin to perform this action"})
			c.Abort()
			return
		}
		//Trim bearer from token
		var tokenString string
		if strings.HasPrefix(authHeader, "Bearer ") {
			tokenString = strings.TrimPrefix(authHeader, "Bearer ")
		} else {
			tokenString = authHeader
		}
		//Decode token
		token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(secretKey), nil
		})
		//Take claim from token
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			userId, userIdOk := claims["userId"].(float64)
			roleId, roleIdOk := claims["roleId"].(float64)
			if !roleIdOk {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "This user does not have role yet"})
				c.Abort()
				return
			}
			if !userIdOk {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
				c.Abort()
				return
			}
			//If have claims set into context
			c.Set("userId", int(userId))
			c.Set("roleId", int(roleId))
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		c.Next()
	}
}
