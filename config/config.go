package config

import (
    "github.com/spf13/viper"
    "log"
)

func LoadConfig() {
    viper.SetConfigName("config")
    viper.SetConfigType("yaml")
    viper.AddConfigPath(".")
    if err := viper.ReadInConfig(); err != nil {
        log.Fatalf("Error reading config file, %s", err)
    }
}

func GetConfig(key string) string {
    return viper.GetString(key)
}
