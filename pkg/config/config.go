package config

import (
	"os"
	"path/filepath"
	"runtime"

	"github.com/spf13/viper"
)

func InitConfig() {
	home, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	viper.AddConfigPath(home)
	viper.SetConfigName(".ayah-sender")
	viper.SetConfigType("yaml")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		println("Using config file:", viper.ConfigFileUsed())
	}
}

func GetProjectRoot() string {
	_, b, _, _ := runtime.Caller(0)
	return filepath.Join(filepath.Dir(b), "../..")
}

func GetDataDir() string {
	return filepath.Join(GetProjectRoot(), "data")
}
