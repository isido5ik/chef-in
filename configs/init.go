package configs

import "github.com/spf13/viper"

func InitConfigs() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
