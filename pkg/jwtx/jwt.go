package jwtx

import (
	"errors"
	"github.com/golang-jwt/jwt"
	"time"
)

type User struct {
	Email string `json:"email"`
}
type MyClaims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}
const TokenExpireTime = time.Hour * 2
var JwtSecret  = []byte("sdfsf!@#SDfsd3423")
//生成token
func MakeToken(email string)(string,error){
	claim := MyClaims{
		Email:          email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireTime).Unix(),
			Issuer: "dev_test",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,claim)
	return token.SignedString(JwtSecret)
}
//解析token
func ParseToken(tokenString string)(*MyClaims, error){
	// 后面是一个匿名函数
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return JwtSecret, nil
	})
	if err != nil {
		return nil, err
	}
	// 校验token
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}