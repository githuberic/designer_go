package e3

import "fmt"

type Mp4Player struct {
}

func (s *Mp4Player) PlayVlc(fileName string) {
}

func (s *Mp4Player) PlayMp4(fileName string) {
	fmt.Println("Playing mp4 file. Name: " + fileName)
}
