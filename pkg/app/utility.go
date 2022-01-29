package app

const (
	EnvLocal      = "local"
	EnvDev        = "dev"
	EnvTest       = "test"
	EnvProduction = "production"
)

type CheckIF interface {
	Check() error
}

func IsEnvLocal() bool {
	return CurrentEnv() == EnvLocal
}

func IsEnvDev() bool {
	return CurrentEnv() == EnvDev
}

func IsEnvTest() bool {
	return CurrentEnv() == EnvTest
}

func IsEnvProduction() bool {
	return CurrentEnv() == EnvProduction
}

func CurrentEnv() string {
	if CfgInstance != nil {
		return CfgInstance.General.Env
	}
	return ""
}
