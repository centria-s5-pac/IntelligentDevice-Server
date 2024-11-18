package common

import (
	"log"
	"sync"

	"github.com/spf13/viper"
)

var once sync.Once

func initConfig() {
	once.Do(func() {
		viper.SetConfigName("config")
		viper.AddConfigPath(".")
		viper.AutomaticEnv()
		if err := viper.ReadInConfig(); err != nil {
			log.Fatalf("Error reading config file, %s", err)
		}
	})
}

func GetConfigString(key string) string {
	initConfig()

	return viper.GetString(key)
}
