package libs

import (
	"fmt"
	"strings"
	"time"

	"github.com/bveranoc/mu_server/pkg/config"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

var JWT_SECRET = []byte(config.GetEnv("JWT_SECRET"))

func GenerateToken(userId uint) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": userId,
		"exp": time.Now().Add(time.Hour * 24 * 7).Unix(),
	})
	tokenStr, err := token.SignedString(JWT_SECRET)
	if err != nil {
		return ""
	}
	return tokenStr
}

func parseAuthorizationHeader(c echo.Context) string {
	authHeader := c.Request().Header["Authorization"]
	return strings.Split(authHeader[0], " ")[1]
}

func parseToken(c echo.Context) (jwt.MapClaims, error) {
	tokenString := parseAuthorizationHeader(c)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signint method: %v", token.Header["alg"])
		}
		return JWT_SECRET, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, err
}

func GetUserID(c echo.Context) (interface{}, error) {
	claims, err := parseToken(c)
	if err != nil {
		return 0, err
	}
	return claims["sub"], nil
}
