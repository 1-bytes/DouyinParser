package auth

import (
	"DouyinParser/pkg/config"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

var jwtSecret []byte

type CustomClaims struct {
	Uid   string `json:"uid"`
	Email string `json:"email"`
}

type Claims struct {
	CustomClaims
	jwt.RegisteredClaims
}

// GenerateToken 创建一个 Token 用于验证.
func GenerateToken(cc CustomClaims) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(3 * time.Hour)
	claims := Claims{
		CustomClaims: cc,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    config.GetString("app.name"),
			ExpiresAt: jwt.NewNumericDate(expireTime),
			IssuedAt:  jwt.NewNumericDate(nowTime),
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	return tokenClaims.SignedString(jwtSecret)
}

// ParseToken 验证 Token.
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	jwt.TimeFunc = time.Now
	return nil, err
}

// CheckToken 检查 Token 是否合法.
func CheckToken(token string) bool {
	_, err := ParseToken(token)
	return err != nil
}
