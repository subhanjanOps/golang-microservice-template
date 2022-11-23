package env

type IEnv interface {
	LoadEnv()
	GetEnvKey(k string) string
	SetEnvKey(k string, val string)
}

type Env struct {
	configPath string
	configName string
}

func NewEnv(configPath string, configName string) *Env {
	return &Env{configPath: configPath, configName: configName}
}
