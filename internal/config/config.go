package config

import (
	"github.com/spf13/viper"
	"github.com/mmagr/planets/internal/model"
	"log"
	"strings"
)

var config = viper.New()

func Init() *viper.Viper {
	config.SetConfigName("config")
	config.AddConfigPath(".")
	config.AddConfigPath("./config")

	if err := config.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			// Config file was found but another error was produced
			log.Fatalf("Error reading config file: %s", err)
		}
	}

	setConfigDefaults()
	return config
}

func setConfigDefaults() {
	config.SetDefault("server.host", "0.0.0.0:8080")
	config.SetDefault("meta.host", "0.0.0.0:8081")
	config.SetDefault("log.level", "info")

	config.SetDefault("planet.vulcano.omega", -5)
	config.SetDefault("planet.vulcano.orbit", 1000)
	config.SetDefault("planet.vulcano.alpha", 0)

	config.SetDefault("planet.ferengi.omega", 1)
	config.SetDefault("planet.ferengi.orbit", 500)
	config.SetDefault("planet.ferengi.alpha", 0)

	config.SetDefault("planet.betasoide.omega", 3)
	config.SetDefault("planet.betasoide.orbit", 2000)
	config.SetDefault("planet.betasoide.alpha", 0)

	config.SetEnvKeyReplacer(strings.NewReplacer("-", "_", ".", "_"))
	config.AutomaticEnv()
}

func Config() *viper.Viper {
	return config
}

func Planet(v *viper.Viper) model.Planet {
	return model.Planet{
		Omega: v.GetFloat64("omega"),
		Alpha: v.GetFloat64("alpha"),
		Orbit: v.GetFloat64("orbit"),
	}
}