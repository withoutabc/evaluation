package secret

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type MyClaims struct {
	ID        int64
	Role      string    //角色，详见roleAuth.go
	LoginTime time.Time //登录时间
	Type      string    //token类型
	jwt.StandardClaims
}

var Secret = []byte("YJX")

func keyFunc(*jwt.Token) (i interface{}, err error) {
	return Secret, nil
}

// GenToken GenToken生成aToken和rToken
func GenToken(userId int64, role string) (aToken, rToken string, c MyClaims, err error) {
	// 创建一个我们自己的声明
	c = MyClaims{
		ID:        userId,
		Role:      role,
		LoginTime: time.Now(),
		Type:      "access",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 30).Unix(), // 过期时间
			Issuer:    "YJX",                                   // 签发人
		},
	}
	// 加密并获得完整的编码后的字符串token
	aToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, c).SignedString(Secret)
	// refresh token 不需要存任何自定义数据
	rToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, MyClaims{
		ID:        userId,
		Role:      role,
		LoginTime: time.Now(),
		Type:      "refresh",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), // 过期时间
			Issuer:    "YJX",                                 // 签发人
		},
	}).SignedString(Secret)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	return
}

// RefreshToken 刷新aToken
func RefreshToken(rToken string) (newAToken, newRToken string, c MyClaims, err error) {
	var claims MyClaims
	_, err = jwt.ParseWithClaims(rToken, &claims, keyFunc)
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return "", "", MyClaims{}, errors.New("invalid refresh token signature")
		}
		return "", "", MyClaims{}, err
	}
	//判断类型是否正确
	if claims.Type != "refresh" {
		return "", "", MyClaims{}, errors.New("错误的类型")
	}
	//生成新的token
	newAToken, newRToken, c, err = GenToken(claims.ID, claims.Role)
	if err != nil {
		return "", "", MyClaims{}, err
	}
	return newAToken, newRToken, c, nil
}

// ParseToken 解析JWT
func ParseToken(tokenString string) (*MyClaims, error) {
	// 解析token
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, keyFunc)
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid { // 校验token
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
