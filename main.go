package main

import (
	"fmt"
	"os"
	"post/command"
	"post/common/logger"
	"post/config"

	"github.com/spf13/viper"
)

func readConfig() {
	var err error

	viper.SetConfigFile("base.env")
	viper.SetConfigType("props")
	err = viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
		return
	}

	if _, err := os.Stat("base.env"); os.IsNotExist(err) {
		fmt.Println("WARNING: file base.env not found")
	} else {
		viper.SetConfigFile("base.env")
		viper.SetConfigType("props")
		err = viper.MergeInConfig()
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	err = viper.Unmarshal(&config.Config)
	for _, key := range viper.AllKeys() {
		viper.BindEnv(key)
	}
}

func main() {
	readConfig()
	raventClient := logger.NewRavenClient()
	logger.NewLogger(raventClient)
	command.Execute()
}
