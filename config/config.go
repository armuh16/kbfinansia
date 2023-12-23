package config

import (
	"fmt"
	"github.com/spf13/viper"
	str2duration "github.com/xhit/go-str2duration/v2"
	"log"
	"time"
)

var c *Config

func Get() *Config {
	return c
}

const (
	Development = "Development"
	Production  = "Production"
)

type Config struct {
	Environment string `mapstructure:"ENV"`
	Port        string `mapstructure:"PORT"`

	// Connection Database Env
	DBUser     string `mapstructure:"DB_USER"`
	DBPassword string `mapstructure:"DB_PASSWORD"`
	DBHost     string `mapstructure:"DB_HOST"`
	DBPort     string `mapstructure:"DB_PORT"`
	DBName     string `mapstructure:"DB_NAME"`

	// Token
	ExpireAccessToken          string `mapstructure:"AUTH_EXPIRE_ACCESS_TOKEN"`
	ExpireRefreshToken         string `mapstructure:"AUTH_EXPIRE_REFRESH_TOKEN"`
	Secret                     string `mapstructure:"SECRET"`
	ExpireAccessTokenDuration  time.Duration
	ExpireRefreshTokenDuration time.Duration
}

func SetConfig() {
	var err error

	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.SetConfigName("config.env")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Cannot read config file .env : ", err.Error())
	}

	if err = viper.Unmarshal(&c); err != nil {
		log.Fatal("Environment can't be loaded: ", err.Error())
	}

	if c.Environment == "development" {
		log.Println("The App is running in development env")
	}

	c.ExpireAccessTokenDuration, err = str2duration.ParseDuration(c.ExpireAccessToken)
	if err != nil {
		panic(fmt.Sprintf("config auth access expired duration string not valid: %s", err.Error()))
	}

	c.ExpireRefreshTokenDuration, err = str2duration.ParseDuration(c.ExpireRefreshToken)
	if err != nil {
		panic(fmt.Sprintf("config auth refresh expired duration string not valid: %s", err.Error()))
	}

	viper.WatchConfig()
}
