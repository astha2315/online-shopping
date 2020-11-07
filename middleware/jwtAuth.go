package middleware

import (
	"fmt"
	"net/http"
	"online-shopping/service"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthorizeJWT(next gin.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		const BEARER_SCHEMA = "Bearer"
		authHeader := c.GetHeader("Authorization")
		tokenString := authHeader[len(BEARER_SCHEMA):]
		token, err := service.JWTAuthService().ValidateToken(tokenString)
		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			fmt.Println(claims)
			next(c)
		} else {
			fmt.Println(err)
			c.AbortWithStatus(http.StatusUnauthorized)
		}

	}
}

func GetJwtClaimsFromRequest(c *gin.Context) jwt.MapClaims {

	// user := c.Get(middleware.DefaultJWTConfig.ContextKey).(*jwt.Token)
	// return user.Claims.(*JwtClaims)
	const BEARER_SCHEMA = "Bearer"
	authHeader := c.GetHeader("Authorization")

	tokenString := strings.TrimSpace(authHeader[len(BEARER_SCHEMA):])
	token, _ := service.JWTAuthService().ValidateToken(strings.TrimSpace(tokenString))

	claims := token.Claims.(jwt.MapClaims)

	return claims
}
