package e1

import "fmt"

type ExistPlayer struct {
}

func (*ExistPlayer) playMp3(fileName string) {
	fmt.Println("play mp3:", fileName)
}

func (*ExistPlayer) playWma(fileName string) {
	fmt.Println("play wma:", fileName)
}
