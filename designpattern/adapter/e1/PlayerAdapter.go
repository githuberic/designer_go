package e1

import "fmt"

type PlayerAdapter struct {
	player ExistPlayer
}

func (this *PlayerAdapter) play(fileType string, fileName string) {
	switch fileType {
	case "mp3":
		this.player.playMp3(fileName)
	case "wma":
		this.player.playMp3(fileName)
	default:
		fmt.Println("Not supported")
	}
}
