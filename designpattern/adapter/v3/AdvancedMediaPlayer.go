package v3

type AdvancedMediaPlayer interface {
	PlayVlc(fileName string)
	PlayMp4(fileName string)
}
