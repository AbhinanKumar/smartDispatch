package midddleware

import (
	"github.com/AbhinanKumar/smart-dispatch/internal/config"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "authorization header missing",
			})
			c.Abort()
			return
		}

		// Extract token
		if !strings.HasPrefix(authHeader, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "invalid token format",
			})
			c.Abort()
			return
		}

		// Validate JWT
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		token, err := config.ValidateJWT(tokenString)

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "invalid token",
			})
			c.Abort()
			return
		}

		claims := token.Claims.(jwt.MapClaims)

		//store UserID
		c.Set("userID", claims["user_id"])
		c.Set("email", claims["email"])
		c.Next()
	}
}
