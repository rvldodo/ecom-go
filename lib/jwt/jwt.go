package jwt

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"github.com/dodo/ecom/config"
)

var (
	defaultMethodJWT = jwt.SigningMethodHS256
	secretTokenKey   = []byte(config.Envs.SecretTokenJWT)
	jwtExpiredTime   = time.Second * time.Duration(config.Envs.JWTExpiredTime)
)

func CreateToken(userID int) (string, error) {
	token := jwt.NewWithClaims(defaultMethodJWT, jwt.MapClaims{
		"userID":    strconv.Itoa(userID),
		"expiredAt": time.Now().Add(jwtExpiredTime).Unix(),
	})

	tokenString, err := token.SignedString(secretTokenKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func CreateTokenCookie(v interface{}) (string, error) {
	token := jwt.NewWithClaims(defaultMethodJWT, jwt.MapClaims{
		"data": v,
	})

	tokenString, err := token.SignedString(secretTokenKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func GetTokenFromRequest(r *http.Request) string {
	token := r.Header.Get("Authorization")
	if token != "" {
		return token
	}

	return ""
}

func ValidateToken(t string) (*jwt.Token, error) {
	return jwt.Parse(t, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}

		return []byte(config.Envs.SecretTokenJWT), nil
	})
}
