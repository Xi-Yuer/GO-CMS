package jwt

import (
	"errors"
	"github.com/Xi-Yuer/cms/config"
	"github.com/Xi-Yuer/cms/responsies"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type Jsonwebtoken struct {
}

func (j *Jsonwebtoken) GenerateTokenUsingHs256(jwtPayload *responsies.JWTPayload) (string, error) {
	claim := responsies.JWTPayload{
		ID:               jwtPayload.ID,
		NickName:         jwtPayload.NickName,
		Role:             jwtPayload.Role,
		PagePermission:   jwtPayload.PagePermission,
		ButtonPermission: jwtPayload.ButtonPermission,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "Xi-Yuer",                                          // 签发者
			Subject:   jwtPayload.NickName,                                // 签发对象
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)), //过期时间
			NotBefore: jwt.NewNumericDate(time.Now().Add(time.Second)),    //最早使用时间
			IssuedAt:  jwt.NewNumericDate(time.Now()),                     //签发时间
			ID:        jwtPayload.ID,                                      // wt ID, 类似于盐值
		},
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claim).SignedString([]byte(config.Config.APP.JWT))
	return token, err
}

func (j *Jsonwebtoken) ParseTokenHs256(tokenStr string) (*responsies.JWTPayload, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &responsies.JWTPayload{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Config.APP.JWT), nil //返回签名密钥
	})
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("claim invalid")
	}

	claims, ok := token.Claims.(*responsies.JWTPayload)
	if !ok {
		return nil, errors.New("invalid claim type")
	}

	return claims, nil
}
