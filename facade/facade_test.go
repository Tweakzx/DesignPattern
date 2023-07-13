package facade

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUserService_LoginOrRegister(t *testing.T) {
	tests := []struct {
		name     string
		username string
		code     string
		want     string
	}{
		{
			name:     "test login",
			username: "user0",
			code:     "123456",
			want:     "user0 login success",
		},
		{
			name:     "test register",
			username: "user1",
			code:     "123456",
			want:     "user1 register success",
		},
	}

	us := &UserService{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := us.LoginOrRegister(tt.username, tt.code)
			require.Nil(t, err)
			assert.Equal(t, tt.want, got.Msg)
		})
	}
}
