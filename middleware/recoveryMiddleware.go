package middleware

import (
	"fmt"
	"net/http"
	"runtime/debug"

	"github.com/gin-gonic/gin"
)

func RecoveryMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": "Internal Server Error",
				})
				fmt.Println("Panic occurred:", err)
				debug.PrintStack() // Cetak stack trace
				c.Abort()
			}
		}()
		c.Next()
	}
}