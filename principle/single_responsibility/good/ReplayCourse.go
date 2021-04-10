package good

import (
	"fmt"
)

type ReplayCourse struct {
	CourseInfo
}

func NewReplayCourse(id int, name string) IGoodCourse {
	return &ReplayCourse{
		CourseInfo{
			iID: id,
			sName: name,
		},
	}
}

func (p *ReplayCourse) Controller() IPlayControl {
	return p
}
func (p *ReplayCourse) Play() {
	fmt.Printf("%v play\n", p.Name())
}
func (p *ReplayCourse) Pause() {
	fmt.Printf("%v pause\n", p.Name())
}
func (p *ReplayCourse) Forward(seconds int) {
	fmt.Printf("%v forward %v\n", p.Name(), seconds)
}
func (p *ReplayCourse) Backward(seconds int) {
	fmt.Printf("%v backward %v\n", p.Name(), seconds)
}