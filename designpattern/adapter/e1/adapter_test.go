package e1

import "testing"

func TestAdapterDP(t *testing.T) {
	player := PlayerAdapter{}
	player.play("mp3", "死了都要爱")
	player.play("wma", "滴滴")
	player.play("mp4", "复仇者联盟")
}
