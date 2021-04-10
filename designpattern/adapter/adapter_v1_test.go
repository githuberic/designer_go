package adapter

import (
	"fmt"
	"testing"
)

type MediaPlayer interface {
	Play(audioType string, fileName string)
}

type AdvancedMediaPlayer interface {
	PlayVlc(fileName string)
	PlayMp4(fileName string)
}

type VlcPlayer struct {
}
func (s *VlcPlayer) PlayVlc(fileName string) {
	fmt.Print("Playing vlc file. Name: " + fileName)
}
func (s *VlcPlayer) PlayMp4(fileName string) {
}

type Mp4Player struct {
}
func (s *Mp4Player) PlayVlc(fileName string) {
}
func (s *Mp4Player) PlayMp4(fileName string) {
	fmt.Println("Playing mp4 file. Name: " + fileName)
}

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

type AudioPlayer struct {
	mediaAdapter MediaAdapter
}

func (s *AudioPlayer) Play(audioType string, fileName string) {
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

func TestVerify(t *testing.T)  {
	var audioType string = "mp4"
	var fileName string = "./old.mp4"

	var audioPlayer = AudioPlayer{}
	audioPlayer.Play(audioType,fileName)
}