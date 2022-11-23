package jwt

import (
	"app-server-gateway-service/pkg/custom_error"
	"github.com/golang-jwt/jwt/v4"
	"os"
	"time"
)

func (a *AuthToken) GetAuthToken() *string {
	token := jwt.New(a.signedMethod)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(a.expiredAfter)
	claims["authorized"] = true
	claims["user_id"] = a.userId
	claims["username"] = a.username
	tokenStr, err := token.SignedString(a.secret)
	customError.CheckError(err)
	return &tokenStr
}

func NewAuthToken(expiredAfter time.Duration, userId int32, username string) *AuthToken {
	return &AuthToken{expiredAfter: expiredAfter, userId: userId, username: username, secret: os.Getenv("JWT_SECRET"), signedMethod: jwt.SigningMethodEdDSA}
}
