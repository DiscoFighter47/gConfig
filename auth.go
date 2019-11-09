package gconfig

import (
	"log"
	"sync"
	"time"

	gero "github.com/DiscoFighter47/gEro"
	"github.com/spf13/viper"
)

// AuthConf ..
type AuthConf struct {
	Secret             string
	TokenExpireTimeout time.Duration
}

func (auth *AuthConf) validate() {
	errV := gero.ValidationError{}
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

var auth *AuthConf
var authOnce sync.Once

func loadAuth() {
	auth = &AuthConf{
		Secret:             viper.GetString("auth.secret"),
		TokenExpireTimeout: viper.GetDuration("auth.token_expire_timeout") * time.Minute,
	}
}

// Auth ...
func Auth() *AuthConf {
	authOnce.Do(func() {
		loadAuth()
		auth.validate()
	})
	return auth
}
