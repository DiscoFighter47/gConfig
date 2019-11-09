package gconfig_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	gconfig "github.com/DiscoFighter47/gConfig"
	"github.com/spf13/viper"
)

func TestAuth(t *testing.T) {
	t.Run("read auth", func(t *testing.T) {
		viper.Set("auth.secret", "secret")
		viper.Set("auth.token_expire_timeout", 1)
		auth := gconfig.Auth()
		assert.Equal(t, "secret", auth.Secret)
		assert.Equal(t, 1*time.Minute, auth.TokenExpireTimeout)

		auth2 := gconfig.Auth()
		assert.Same(t, auth, auth2)
	})
}
