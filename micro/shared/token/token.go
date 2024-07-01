package token

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

// 定义一个密钥，用于签名和验证 Token
var jwtKey = []byte("12345699999")

// Claims 结构定义了 Token 中的声明
type Claims struct {
	UserID   int    `json:"user_id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

// CreateToken 创建 Token
func CreateToken(userID int, username string) (string, error) {
	// 设置 Token 的过期时间
	expirationTime := time.Now().Add(5 * time.Minute)

	// 创建声明
	claims := &Claims{
		UserID:   userID,
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// 创建 Token 对象并签名
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

// VerifyToken 验证
func VerifyToken(tokenString string) (*Claims, error) {
	// 解析 Token
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		return nil, err
	}

	// 检查 Token 的有效性
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, fmt.Errorf("invalid token")
	}
}
