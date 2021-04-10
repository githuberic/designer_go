package good

import "fmt"

type JavaCourse struct {
	iID          int
	sName        string
	xCurrentUser IUser
}

func NewJavaCourse() ICourse {
	return &JavaCourse{
		iID:          12,
		sName:        "Java",
		xCurrentUser: nil,
	}
}
func (p *JavaCourse) ID() int {
	return p.iID
}
func (p *JavaCourse) Name() string {
	return p.sName
}
func (p *JavaCourse) SetUser(user IUser) {
	p.xCurrentUser = user
}
func (p *JavaCourse) Study() {
	fmt.Printf("%v is learning %v\n", p.xCurrentUser.Name(), p.Name())
}
