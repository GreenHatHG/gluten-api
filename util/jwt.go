package util

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gluten/global"
	"time"
)

func GetJWTToken(id string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  id,
		"exp": time.Now().Add(time.Hour * time.Duration(global.JwtConfig.Exp)).Unix(),
	})
	if tokenString, err := token.SignedString([]byte(global.JwtConfig.Secret)); err != nil {
		return "", err
	} else {
		return tokenString, nil
	}
}

func ParseToken(tokenString string) (string, bool) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(global.JwtConfig.Secret), nil
	})
	if err != nil {
		Logger.Error(err)
		return "", false
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims["id"].(string), true
	} else {
		return "", false
	}
}

func GetJwtId(c *gin.Context) string {
	id, _ := c.Get("id")
	return id.(string)
}
