package jwtUtil
import (
	jwt "github.com/dgrijalva/jwt-go"
)


type JWT struct {
    SigningKey []byte
}

type CustomClaims struct {
    Name  string `json:"username"`
    Password string `json:"password"`
    jwt.StandardClaims
}


func (j *JWT) CreateToken(claims CustomClaims) (string, error) {
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(j.SigningKey)
}

j := &jwt.JWT{
	[]byte("newtrekWang"),
}


// token, err := j.CreateToken(claims)