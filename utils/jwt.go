package utils

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

var jwtSecret = []byte("Foxconn")

type Claims struct {
	ID        uint   `json:"id;omitempty"`
	UserName  string `json:"user_name;omitempty"`
	Email     string `json:"email;omitempty"`
	Status    string `json:"status;omitempty"`
	Authority string `json:"authority;omitempty"`
	jwt.StandardClaims
}

// 签发token
func GenerateToken(id uint, userName string, email string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(2 * time.Hour)
	claims := Claims{
		ID:       id,
		UserName: userName,
		Email:    email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "Foxconn",
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)
	return token, err
}

// 验证用户token

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err

}
