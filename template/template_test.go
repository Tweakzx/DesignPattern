package template

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSms_Send(t *testing.T) {
	us := NewUnicomSms()
	err := us.Send("hello world", 1237654)
	assert.Nil(t, err)
}
