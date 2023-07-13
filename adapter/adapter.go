package adapter

type Player interface {
	Play() string
}

type MP3Player struct{}

func (m MP3Player) Play() string {
	return "Playing MP3 file"
}

type WAVPlayer struct{}

func (w WAVPlayer) PlayWAV() string {
	return "Playing WAV file"
}

type WAVPlayerAdapter struct {
	wavPlayer WAVPlayer
}

func (a WAVPlayerAdapter) Play() string {
	return a.wavPlayer.PlayWAV()
}
