package env

import (
	"github.com/spf13/viper"
	"log"
)

type Env struct {
	RPC_URL          string `mapstructure:"RPC_URL"`
	PRIVATE_KEY      string `mapstructure:"PRIVATE_KEY"`
	CONTRACT_ADDRESS string `mapstructure:"CONTRACT_ADDRESS"`
}

func NewEnv() *Env {
	env := Env{}
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Can't find the file .env : ", err)
	}

	err = viper.Unmarshal(&env)
	if err != nil {
		log.Fatal("Environment can't be loaded: ", err)
	}

	return &env
}
