package good

import "fmt"

type GolangCourse struct {
	iID          int
	sName        string
	xCurrentUser IUser
}

func NewGolangCourse() ICourse {
	return &GolangCourse{
		iID:          11,
		sName:        "golang",
		xCurrentUser: nil,
	}
}
func (p *GolangCourse) ID() int {
	return p.iID
}
func (p *GolangCourse) Name() string {
	return p.sName
}
func (p *GolangCourse) SetUser(user IUser) {
	p.xCurrentUser = user
}
func (p *GolangCourse) Study() {
	fmt.Printf("%v is learning %v\n", p.xCurrentUser.Name(), p.Name())
}
