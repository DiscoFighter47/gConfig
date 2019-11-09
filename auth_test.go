package gconfig

import (
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

type spyLogger struct {
	fatal string
}

func (l *spyLogger) Fatal(v ...interface{}) {
	for _, x := range v {
		l.fatal += fmt.Sprintf("%v", x)
	}
}

func TestAuth(t *testing.T) {
	t.Run("config error", func(t *testing.T) {
		l := &spyLogger{}
		SetLogger(l)
		Auth()
		assert.Equal(t, "Auth Configuration Error:{\"auth.secret\":[\"missing\"],\"auth.token_expire_timeout\":[\"missing\"]}", l.fatal)
	})

	t.Run("read auth", func(t *testing.T) {
		viper.Set("auth.secret", "secret")
		viper.Set("auth.token_expire_timeout", 1)
		authOnce = new(sync.Once)
		auth := Auth()
		assert.Equal(t, "secret", auth.Secret)
		assert.Equal(t, 1*time.Minute, auth.TokenExpireTimeout)

		auth2 := Auth()
		assert.Same(t, auth, auth2)
	})
}
