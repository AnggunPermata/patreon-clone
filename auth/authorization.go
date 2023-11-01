package auth

import (
	"fmt"
	"time"

	"github.com/anggunpermata/patreon-clone/configs"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func CreateToken(c echo.Context, secretJWT string, userID int, role string, username string, expHour int) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["userId"] = int(userID)
	claims["role"] = role
	claims["username"] = username
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(expHour)).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(configs.SecretJWT))
}

func ExtractTokenUserId(c echo.Context) (int, string) {
	user := c.Get("user").(*jwt.Token)
	if user.Valid {
		claims := user.Claims.(jwt.MapClaims)
		userId := int(claims["userId"].(float64))
		role := fmt.Sprintf("%v", claims["role"])
		return userId, role
	}

	return 0, "invalid token"
}

func IdentifyUserUsingJWTToken(c echo.Context, cfg *configs.Config, tokenStr string) (role string, username string, authorized bool){
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return []byte(cfg.SecretJWT), nil
	})

	if err != nil {
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		role = fmt.Sprintf("%v", claims["role"])
		authorized = claims["authorized"].(bool)
		username = claims["username"].(string)
		return
	}

	return
}