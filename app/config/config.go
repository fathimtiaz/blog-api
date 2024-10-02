package config

import (
	"log"
	"path"
	"runtime"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	App struct {
		Env  Env `envconfig:"APP_ENV" required:"true"`
		Port Env `envconfig:"APP_PORT" required:"true"`
	}
	RateLimit struct {
		IP      Env `envconfig:"RATELIMIT_IP" required:"true"`
		IPRoute Env `envconfig:"RATELIMIT_IPROUTE" required:"true"`
	}
	DB struct {
		ConnStr Env `envconfig:"DB_CONNSTR" required:"true"`
	}
}

// LoadDefault loads default config (default.yml) and override config with env if supplied
func LoadDefault() *Config {
	return load(".env")
}

// load config and populate to config struct
func load(env string) *Config {
	var config Config

	readEnv(env)
	err := envconfig.Process("", &config)
	if err != nil {
		panic(err)
	}
	return &config
}

func readEnv(env string) {
	err := godotenv.Overload(getSourcePath() + "/../" + env)
	if err != nil {
		log.Print(err)
	}
}

func getSourcePath() string {
	_, filename, _, _ := runtime.Caller(1)
	return path.Dir(filename)
}
