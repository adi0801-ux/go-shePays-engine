package utils

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Config struct {
	DSN           string `mapstructure:"DSN"`
	GINMode       string `mapstructure:"GIN_MODE"`
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
	AppId         string `mapstructure:"APP_ID"`
	AppKey        string `mapstructure:"APP_KEY"`
	AppSecret     string `mapstructure:"APP_SECRET"`
	HappyUrl      string `mapstructure:"HAPPY_URL"`
}

func LoadConfig(path string) (config *Config, err error) {
	v := viper.New()

	v.AddConfigPath(path)
	v.SetConfigName("app")
	v.SetConfigType("env") // json, xml

	v.AutomaticEnv()

	err = v.ReadInConfig()
	if err != nil {
		return
	}
	err = v.Unmarshal(&config)

	v.OnConfigChange(func(e fsnotify.Event) {
		Log.Warn("Config file changed:", e.Name)
		err := ReadConfig(v, config)
		if err != nil {
			Log.Error(err)
			return
		}
	})
	v.WatchConfig()

	return
}

func ReadConfig(v *viper.Viper, config *Config) error {
	err := v.ReadInConfig()
	if err != nil {
		return err
	}
	err = v.Unmarshal(&config)
	if err != nil {
		return err
	}

	return nil
}
