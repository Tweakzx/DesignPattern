package adapter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAliyunClientAdapter_CreateServer(t *testing.T) {
	tests := []struct {
		name   string
		player Player
		want   string
	}{
		{"wav", &WAVPlayerAdapter{}, "Playing WAV file"},
		{"mp3", &MP3Player{}, "Playing MP3 file"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.player.Play())
		})
	}
}
