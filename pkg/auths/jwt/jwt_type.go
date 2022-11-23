package jwt

import (
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type IAuthToken interface {
	GetAuthToken() *string
}
type AuthToken struct {
	expiredAfter time.Duration
	userId       int32
	username     string
	secret       string
	signedMethod jwt.SigningMethod
}
