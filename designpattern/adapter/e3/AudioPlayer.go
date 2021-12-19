package e3

import "fmt"

type AudioPlayer struct {
	mediaAdapter MediaAdapter
}

func (s *AudioPlayer) Play(audioType string, fileName string) {
	//播放 mp3 音乐文件的内置支持
	if audioType == ("mp3") {
		fmt.Println("Playing mp3 file. Name: " + fileName)
	} else if audioType == "vlc" || audioType == "mp4" {
		s.mediaAdapter = MediaAdapter{}
		s.mediaAdapter.MediaAdapter(audioType)
		s.mediaAdapter.Play(audioType, fileName)
	} else {
		fmt.Println("Invalid media. " + audioType + " format not supported")
	}
}
