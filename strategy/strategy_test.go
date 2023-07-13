package strategy

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestStrategy_Save(t *testing.T) {
	tests := []struct {
		name  string
		key   string
		value []byte
		want  string
	}{
		{"file", "key1", []byte("value1"), "key1:value1 saved data to file"},
		{"db", "key2", []byte("value2"), "key2:value2 saved data to db"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s, err := NewStorageStrategy(tt.name)
			require.Nil(t, err)
			assert.Equal(t, tt.want, s.Save(tt.key, tt.value))
		})
	}
}
