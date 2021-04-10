package good

import (
	"fmt"
)

type LiveCourse struct {
	CourseInfo
}

func NewLiveCourse(id int, name string) IGoodCourse {
	return &LiveCourse{
		CourseInfo{
			iID:   id,
			sName: name,
		},
	}
}
func (p *LiveCourse) Controller() IPlayControl {
	return p
}
func (p *LiveCourse) Play() {
	fmt.Printf("%v play\n", p.Name())
}
func (p *LiveCourse) Pause() {
	fmt.Printf("%v pause\n", p.Name())
}
