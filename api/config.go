package api

import "github.com/spf13/viper"

type Config struct {
	ADO_URL   string `mapstructure:"ADO_URL"`
	ADO_TOKEN string `mapstructure:"ADO_TOKEN"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("greetctl")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
