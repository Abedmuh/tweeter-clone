package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthenticated"})
			return
		}

		secretKey := viper.GetString("JWT_SECRET_KEY")
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
				return []byte(secretKey), nil
		})
		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{"error": "Unauthenticated"})
		}
		fmt.Println(token)
		
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			c.Set("user", claims["user"])
		} else {
			c.AbortWithStatusJSON(401, gin.H{"error": "fail to get user"})
		}

		c.Next()
	}
}
