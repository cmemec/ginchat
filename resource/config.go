package resource

import "github.com/spf13/viper"

func InitConfig() error {
	viper.SetConfigName("app")
	viper.SetConfigType("yml")
	viper.AddConfigPath("./config")
	return viper.ReadInConfig()
}
