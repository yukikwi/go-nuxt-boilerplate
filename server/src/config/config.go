package config

import (
	"github.com/yukikwi/go-nuxt-boilerplate/utils"
)

type (
	ConfigInterface struct {
		Db          string
		Port        string
		Environment string
	}
)

var Config ConfigInterface

func init() {
	Config = ConfigInterface{
		Db:          utils.GetEnv("DB_DSN", ""),
		Port:        utils.GetEnv("PORT", "3000"),
		Environment: utils.GetEnv("ENVIRONMENT", "development"),
	}
}
