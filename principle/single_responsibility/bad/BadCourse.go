package bad

import "fmt"

type IBadCourse interface {
	ID() int
	Name() string

	Play()
	Pause()
	Forward(int)
	Backward(int)
}

type BadCourse struct {
	iID   int
	sName string
}

func NewBadCourse(id int, name string) *BadCourse {
	return &BadCourse{
		iID:   id,
		sName: name,
	}
}

func (p *BadCourse) ID() int {
	return p.iID
}
func (me *BadCourse) Name() string {
	return me.sName
}

func (me *BadCourse) Play() {
	fmt.Printf("%v play\n", me.Name())
}

func (me *BadCourse) Pause() {
	fmt.Printf("%v pause\n", me.Name())
}

func (p *BadCourse) Forward(seconds int) {
	if p.Name() == "录播课" {
		fmt.Printf("%v forward %v seconds\n", p.Name(), seconds)
	} else {
		fmt.Printf("%v cannot forward\n", p.Name())
	}
}

func (p *BadCourse) Backward(seconds int) {
	if p.Name() == "录播课" {
		fmt.Printf("%v backward %v seconds\n", p.Name(), seconds)
	} else {
		fmt.Printf("%v cannot backward\n", p.Name())
	}
}
