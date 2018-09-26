package config

import (
	"github.com/spf13/viper"
	"strings"
)

type Config interface {
	GetString(key string) string
	GetInt(key string) int
	GetBool(key string) bool
	Init()
}

type viperConfig struct{

}

func (v *viperConfig) Init(){
	viper.SetEnvPrefix("d2d-")
	viper.AutomaticEnv()
	replacer := strings.NewReplacer(`.`,`_`)
	viper.SetEnvKeyReplacer(replacer)
	viper.SetConfigType(`json`)
	viper.SetConfigName(`Configs`)
	viper.AddConfigPath("config/")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func NewViperConfig() Config{
	v := &viperConfig{}
	v.Init()
	return v
}


func (v *viperConfig) GetString(key string) string {
	return viper.GetString(key)
}

func (v *viperConfig) GetInt(key string) int {
	return viper.GetInt(key)
}

func (v *viperConfig) GetBool(key string) bool{
	return viper.GetBool(key)
}
