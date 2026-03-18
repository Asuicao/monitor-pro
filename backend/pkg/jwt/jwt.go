package jwt

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func GetUserID(c *gin.Context) (uint, error) {
	return 1, nil
}

func GenerateToken(userID uint) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
	})
	tokenString, _ := token.SignedString([]byte("secret"))
	return tokenString
}
