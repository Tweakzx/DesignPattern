package proxy

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUserProxy_Login(t *testing.T) {
	proxy := NewUserProxy(&User{})
	err := proxy.Login("admin", "123456")
	require.Nil(t, err)
}
