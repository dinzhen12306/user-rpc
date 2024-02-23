package SDK

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
	"userrpc/config"
)

// @secretKey: JWT 加解密密钥
// @iat: 时间戳
// @seconds: 过期时间，单位秒
// @payload: 数据载体
func GetJwtToken(iat, seconds int64, payload string) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["payload"] = payload
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(config.ServerConfigs.Token.SecretKey))
}

func CheckToken(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		//hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(config.ServerConfigs.Token.SecretKey), nil
	})
	if err != nil {
		return "", err
	}
	claims, _ := token.Claims.(jwt.MapClaims)
	if int64(claims["exp"].(float64)) < time.Now().Unix() {
		return "", err
	}
	return claims["payload"].(string), nil
}
