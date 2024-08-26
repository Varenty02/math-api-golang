package middleware

import (
	"code-math/component/tokenprovider"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)
func extractTokenFromHeaderString(s string) (string, error) {
	parts := strings.Split(s, " ")
	if parts[0] != "Bearer" || len(parts) < 2 || strings.TrimSpace(parts[1]) == "" {
		return "", fmt.Errorf("error cannot extract")
	}
	return parts[1], nil

}
func AuthMiddleware(secretKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString,err := extractTokenFromHeaderString(c.GetHeader("Authorization"))
		if err!=nil{
			panic(err) 
		}
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "authorization header missing"})
			c.Abort()
			return
		}

		claims, err := tokenprovider.ParseToken(tokenString, secretKey)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			c.Abort()
			return
		}

		c.Set("userID", claims.UserID)
		c.Next()
	}
}