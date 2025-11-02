package config

import (
	"github.com/yukikwi/go-nuxt-boilerplate/utils"
)

var Config = map[string]string{
	"Db":   utils.GetEnv("DB_DSN", ""),
	"Port": utils.GetEnv("PORT", "3000"),
}
