package env

import (
	"app-server-gateway-service/pkg/custom_error"
	"errors"
	"github.com/spf13/viper"
	"os"
)

func (e *Env) LoadEnv() {
	vp := viper.New()
	vp.SetConfigName(e.configName)
	vp.AddConfigPath(e.configPath)
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if errors.Is(err, viper.ConfigFileNotFoundError{}) {
		customError.CheckError(errors.New("config file not found"))
	}
}

func (e *Env) GetEnvKey(k string) string {
	return os.Getenv(k)
}

func (e *Env) SetEnvKey(k string, val string) {
	err := os.Setenv(k, val)
	customError.CheckError(err)
}
