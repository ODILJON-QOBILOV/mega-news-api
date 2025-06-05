package middleware

import (
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/newsapi/v2/config"
)

func AuthMiddleware(roles ...string) gin.HandlerFunc {
    return func(c *gin.Context) {
        // Extract token from the Authorization header
        authHeader := c.GetHeader("Authorization")
        if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "missing or invalid token"})
            c.Abort()
            return
        }
        tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

        // Parse the token
        token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
            if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
                return nil, jwt.ErrSignatureInvalid
            }
            return config.JwtSecret, nil
        })
        if err != nil {
            log.Printf("Token parse error: %v", err)
            c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
            c.Abort()
            return
        }

        // Extract claims
        claims, ok := token.Claims.(jwt.MapClaims)
        if !ok {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token claims"})
            c.Abort()
            return
        }

        role, roleOk := claims["role"].(string)
        username, userOk := claims["username"].(string)
        userID, userIDOk := claims["user_id"].(float64)
        if !roleOk || !userOk || !userIDOk {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token claims"})
            c.Abort()
            return
        }

        log.Printf("Claims from token: role=%s, username=%s, user_id=%f", role, username, userID)

        // Check if the role is allowed
        allowed := false
        for _, r := range roles {
            if r == role {
                allowed = true
                break
            }
        }
        if !allowed {
            c.JSON(http.StatusForbidden, gin.H{"error": "access denied"})
            c.Abort()
            return
        }

        // Token expiry check (optional)
        if exp, ok := claims["exp"].(float64); ok {
            if time.Now().Unix() > int64(exp) {
                c.JSON(http.StatusUnauthorized, gin.H{"error": "token expired"})
                c.Abort()
                return
            }
        }

        // Set user context
        c.Set("username", username)
        c.Set("role", role)
        c.Set("user_id", int(userID)) // Save user_id in context

        c.Next()
    }
}
