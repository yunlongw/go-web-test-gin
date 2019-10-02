package util

import (
	"github.com/dgrijalva/jwt-go"
	"go-web-test/pkg/setting"
	"time"
)

var jwtSecret = []byte(setting.JwtSecret)

type Claims struct {
	Username string `json:"username"`
	PassWord string `json:"password"`
	jwt.StandardClaims
}


/**
	生成token
 */
func GenerateToken(username, password string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(3 * time.Hour)

	claims := Claims{
		username,
		password,
		jwt.StandardClaims {
			ExpiresAt : expireTime.Unix(),
			Issuer : "gin-blog",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}

/**
验证 token
 */
func ParseToken(token string) (*Claims , error)  {
	tokenClaims , err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{},  error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid{
			return claims, nil
		}
	}

	return nil, err
}