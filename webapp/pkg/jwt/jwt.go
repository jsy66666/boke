package jwt

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

//创建token 解析token

// 定义jwt过期时间
const TokenExpireDuration = time.Hour * 2

var mySecret = []byte("夏天夏天悄悄过去")

type MyClaims struct {
	UserID int64 `json:"user_id"`
	jwt.StandardClaims
}

func keyFunc(_ *jwt.Token) (i interface{}, err error) {
	return mySecret, nil
}

// 生成jwt  生成token
func GenToken(userID int64) (string, error) {
	//创建一个自己声明的数据
	c := MyClaims{
		userID,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(),
			Issuer:    "webapp",
		},
	}
	//使用指定的签名方式创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	//使用指定的secret签名并获得完整的token
	return token.SignedString(mySecret)
}

// 解析jwt  解析token
func ParseToken(tokenstring string) (claim *MyClaims, err error) {

	var token *jwt.Token
	claim = &MyClaims{}
	token, err = jwt.ParseWithClaims(tokenstring, claim, keyFunc)
	if err != nil {
		return
	}
	if !token.Valid { //校验token
		err = errors.New("invalid token")
	}
	return
}
