package bad

import "fmt"

type BadUser struct {
	iID   int
	sName string
}

func NewBadUser(id int, name string) *BadUser {
	return &BadUser{
		iID:   id,
		sName: name,
	}
}

func (p *BadUser) StudyJavaCourse() {
	fmt.Printf("%v is learning %v\n", p.sName, "java")
}
func (p *BadUser) StudyGolangCourse() {
	fmt.Printf("%v is learning %v\n", p.sName, "golang")
}
