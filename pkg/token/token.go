package token

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
)

var tokenKey = []byte("ch6tf2wI39TuHQ74CmAMO9JkNlY8KDNq")

type UserClaims struct {
	Identity string `json:"identity"`
	Email    string `json:"email"`
	jwt.StandardClaims
}

// GenerateToken 生成token
func GenerateToken(identity, email string) (string, error) {
	userClaim := &UserClaims{
		Identity:       identity,
		Email:          email,
		StandardClaims: jwt.StandardClaims{},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, userClaim)
	tokenString, err := token.SignedString(tokenKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// AnalyseToken 解析token
func AnalyseToken(tokenString string) (*UserClaims, error) {
	userClaim := &UserClaims{}
	claims, err := jwt.ParseWithClaims(tokenString, userClaim, func(token *jwt.Token) (interface{}, error) {
		return tokenKey, nil
	})
	if err != nil {
		return nil, err
	}
	if !claims.Valid {
		return nil, errors.New("invalid token")
	}
	return userClaim, nil
}
