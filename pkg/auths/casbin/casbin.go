package casbin

import (
	customError "app-server-gateway-service/pkg/custom_error"
	"github.com/casbin/casbin/v2"
	"log"
)

type IAuthToken interface {
	GetAuthToken() string
}

type AuthToken struct {
}

func (a *AuthToken) GetAuthToken() string {
	//TODO implement me
	enforcer, err := casbin.NewEnforcer("./model.conf", "./policy.csv")
	customError.CheckError(err)
	//enforcer.
	log.Println(enforcer)
	return ""
}

func NewAuthToken() *AuthToken {
	return &AuthToken{}
}
