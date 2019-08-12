package config

import (
	"log"
	"sync"
	"time"

	gson "github.com/DiscoFighter47/gSON"
	"github.com/spf13/viper"
)

// AuthCfg ..
type AuthCfg struct {
	Secret             string
	TokenExpireTimeout time.Duration
}

func (auth *AuthCfg) validate() {
	errV := gson.ValidationError{}
	if auth.Secret == "" {
		errV.Add("auth.secret", "missing")
	}
	if auth.TokenExpireTimeout == 0 {
		errV.Add("auth.token_expire_timeout", "missing")
	}
	if len(errV) > 0 {
		log.Fatal("Auth Configuration Error:", errV)
	}
}

var auth *AuthCfg
var authOnce sync.Once

func loadAuth() {
	auth = &AuthCfg{
		Secret:             viper.GetString("auth.secret"),
		TokenExpireTimeout: viper.GetDuration("auth.token_expire_timeout") * time.Minute,
	}
}

// Auth ...
func Auth() *AuthCfg {
	authOnce.Do(func() {
		loadAuth()
		auth.validate()
	})
	return auth
}
