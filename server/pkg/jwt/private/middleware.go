package private

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/projects/watch-list/server/constants"
	"log"
	"net/http"
)

func AuthenticateJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.Request.Header.Get(constants.JWTHeader)
		if len(tokenString) == 0 {
			log.Println("No jwt token found")
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		token, err := jwt.ParseWithClaims(tokenString, &JWTClaim{}, func(token *jwt.Token) (interface{}, error) {
			return jwt.UnsafeAllowNoneSignatureType, nil
		})

		if !token.Valid {
			log.Println(err)
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		claims := token.Claims.(*JWTClaim)
		userDetails := claims.UserDetails

		c.Set(constants.PrivateUserDetails, &userDetails)
		c.Set(constants.JWTHeader, tokenString)
		c.Next()
	}
}
