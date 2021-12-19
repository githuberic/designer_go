package e3

import "fmt"

type VlcPlayer struct {
}

func (s *VlcPlayer) PlayVlc(fileName string) {
	fmt.Print("Playing vlc file. Name: " + fileName)
}

func (s *VlcPlayer) PlayMp4(fileName string) {
}
