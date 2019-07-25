package gconfig

import (
	"log"

	"github.com/spf13/viper"
	// viper/remote
	_ "github.com/spf13/viper/remote"
)

func init() {
	viper.BindEnv("consul_url")
	viper.BindEnv("consul_path")

	errC := configError{}

	if viper.GetString("consul_url") == "" {
		errC.add("CONSUL_URL", "missing")
	}
	if viper.GetString("consul_path") == "" {
		errC.add("CONSUL_PATH", "missing")
	}

	viper.AddRemoteProvider("consul", viper.GetString("consul_url"), viper.GetString("consul_path"))
	viper.SetConfigType("yml")

	if err := viper.ReadRemoteConfig(); err != nil {
		errC.add("Read Config", err.Error())
	}

	if len(errC) > 0 {
		log.Fatal("Configuration Error:", errC)
	}
}
