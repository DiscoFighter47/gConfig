package gconfig

import (
	gero "github.com/DiscoFighter47/gEro"
	"github.com/spf13/viper"

	// viper/remote
	_ "github.com/spf13/viper/remote"
)

// ReadConfig ...
func ReadConfig() {
	viper.BindEnv("consul_url")
	viper.BindEnv("consul_path")

	errV := gero.ValidationError{}

	if viper.GetString("consul_url") == "" {
		errV.Add("CONSUL_URL", "missing")
	}
	if viper.GetString("consul_path") == "" {
		errV.Add("CONSUL_PATH", "missing")
	}

	viper.AddRemoteProvider("consul", viper.GetString("consul_url"), viper.GetString("consul_path"))
	viper.SetConfigType("yml")

	if err := viper.ReadRemoteConfig(); err != nil {
		errV.Add("Read Config", err.Error())
	}

	if len(errV) > 0 {
		logger.Fatal("Configuration Error:", errV)
	}
}
