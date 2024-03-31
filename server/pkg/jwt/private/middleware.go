package private

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/projects/watch-list/server/constants"
	"log"
	"net/http"
	"os"
	"strings"
)

func AuthenticateJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		bearerToken, err := extractToken(c)
		if err != nil {
			log.Println("unable to extract token from header")
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		token, err := jwt.ParseWithClaims(bearerToken, &JWTClaim{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("SECRET_KEY")), nil
		})
		if err != nil {
			log.Println(err)
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		if !token.Valid {
			log.Println(err)
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		claims := token.Claims.(*JWTClaim)

		c.Set(constants.PrivateUserDetails, &UserDetails{
			UserID: claims.Sub,
			Email:  claims.Email,
		})
		c.Set(constants.JWTHeader, bearerToken)
		c.Next()
	}
}

func extractToken(c *gin.Context) (string, error) {
	bearerToken := c.Request.Header.Get(constants.Authorization)
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1], nil
	}
	return "", errors.New("no auth token found")
}
