package logic

import (
	"context"
	"time"

	"github.com/golang-jwt/jwt"
)

type LogicJwt struct {
}

const (
	ExpiresTime = time.Hour * 1
)

var (
	MyJwt     LogicJwt = LogicJwt{}
	JwtSecret []byte   = []byte("123456")
)

type MyClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// GenerateToken 生成 jwt 格式 token
func (LogicJwt) GenerateToken(ctx context.Context, username string) (token string, err error) {
	tokenHeader := jwt.NewWithClaims(jwt.SigningMethodHS256, &MyClaims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(ExpiresTime).Unix(),
		},
	})
	token, err = tokenHeader.SignedString(JwtSecret)
	return
}

// Valid 验证 token 是否合法
func (LogicJwt) Valid(ctx context.Context, token string) (valid bool) {
	var claims *MyClaims = &MyClaims{}
	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return JwtSecret, nil
	})
	if err != nil {
		valid = false
	} else {
		valid = tkn.Valid
	}
	return
}

// 解析 token 内容
func (LogicJwt) Parse(ctx context.Context, token string) (claims *MyClaims, valid bool) {
	claims = &MyClaims{}
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return JwtSecret, nil
	})
	if err != nil {
		valid = false
	} else {
		valid = true
	}
	return
}
