package e3

type MediaAdapter struct {
	advancedMediaPlayer AdvancedMediaPlayer
}

func (s *MediaAdapter) MediaAdapter(audioType string) {
	if audioType == "vlc" {
		s.advancedMediaPlayer = new(VlcPlayer)
	} else if audioType == "mp4" {
		s.advancedMediaPlayer = new(Mp4Player)
	}
}

func (s *MediaAdapter) Play(audioType string, fileName string) {
	if audioType == "vlc" {
		s.advancedMediaPlayer.PlayVlc(fileName)
	} else if audioType == "mp4" {
		s.advancedMediaPlayer.PlayMp4(fileName)
	}
}
