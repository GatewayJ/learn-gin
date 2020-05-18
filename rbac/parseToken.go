package rbac

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// 一些常量
var (
	TokenExpired     error  = errors.New("Token is expired")
	TokenNotValidYet error  = errors.New("Token not active yet")
	TokenMalformed   error  = errors.New("That's not even a token")
	TokenInvalid     error  = errors.New("Couldn't handle this token:")
	SignKey          string = "newtrekWang"
)

type JWT struct {
	SigningKey []byte
}

type CustomClaims struct {
	ID    string `json:"userId"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
	jwt.StandardClaims
}

//创建token
func (j *JWT) CreateToken(claims CustomClaims) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	res, err := token.SignedString(j.SigningKey)
	fmt.Println("err:", err)
	return res, err
}

//解析token
func (j *JWT) ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			log.Panicln("unexpected signing method")
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return j.SigningKey, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, TokenInvalid
}

//更新token
func (j *JWT) RefreshToken(tokenString string) (string, error) {
	jwt.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}

	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})

	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		jwt.TimeFunc = time.Now
		claims.StandardClaims.ExpiresAt = time.Now().Add(1 * time.Hour).Unix()
		return j.CreateToken(*claims)
	}

	return "", TokenInvalid
}

func GenerateToken(w string) (string, error) {
	j := &JWT{[]byte("man")}
	claims := CustomClaims{
		"1", "Jaya", "123456", jwt.StandardClaims{
			NotBefore: int64(time.Now().Unix() - 1000),
			ExpiresAt: int64(time.Now().Unix() + 3600),
			Issuer:    "man",
		},
	}

	token, err := j.CreateToken(claims)
	return token, err
}

func ParesToken(w string) (string, error) {
	j := &JWT{[]byte("man")}
	claims, err := j.ParseToken(w)
	return claims.Name, err
}
