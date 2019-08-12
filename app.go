package config

import (
	"log"
	"sync"
	"time"

	gson "github.com/DiscoFighter47/gSON"
	"github.com/spf13/viper"
)

// AppCfg ..
type AppCfg struct {
	Port         int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	IdelTimeout  time.Duration
	GraceTimeout time.Duration
}

func (app *AppCfg) validate() {
	errV := gson.ValidationError{}
	if app.Port == 0 {
		errV.Add("app.port", "missing")
	}
	if app.ReadTimeout == 0 {
		errV.Add("app.read_timeout", "missing")
	}
	if app.WriteTimeout == 0 {
		errV.Add("app.write_timeout", "missing")
	}
	if app.IdelTimeout == 0 {
		errV.Add("app.idle_timeout", "missing")
	}
	if app.GraceTimeout == 0 {
		errV.Add("app.grace_timeout", "missing")
	}
	if len(errV) > 0 {
		log.Fatal("App Configuration Error:", errV)
	}
}

var app *AppCfg
var appOnce sync.Once

func loadApp() {
	app = &AppCfg{
		Port:         viper.GetInt("app.port"),
		ReadTimeout:  viper.GetDuration("app.read_timeout") * time.Second,
		WriteTimeout: viper.GetDuration("app.write_timeout") * time.Second,
		IdelTimeout:  viper.GetDuration("app.idle_timeout") * time.Second,
		GraceTimeout: viper.GetDuration("app.grace_timeout") * time.Second,
	}
}

// App ...
func App() *AppCfg {
	appOnce.Do(func() {
		loadApp()
		app.validate()
	})
	return app
}
