package config

import (
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env            string `yaml:"env" env-default:"local"`
	StoragePath    string `yaml:"storage_path" env-required:"true"`
	MigrationsPath string
	GRPC           GRPCConfig    `yaml:"grpc"`
	TokenTTL       time.Duration `yaml:"token_ttl" env-default:"1h"`
}

type GRPCConfig struct {
	Port    int           `yaml:"port"`
	Timeout time.Duration `yaml:"timeout"`
}

func LoadCfg() *Config {
	var cfg Config
	err := cleanenv.ReadConfig("/Users/limon4ik/Desktop/проекты/shop-dn/services/auth/config.yml", &cfg)
	if err != nil {
		panic(err)
	}
	return &cfg
}
