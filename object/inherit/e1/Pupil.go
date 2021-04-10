package e1

import "fmt"

type Pupil struct {
	Student
}

func (p *Pupil) testing()  {
	fmt.Println("小学生正在考试中...")
}