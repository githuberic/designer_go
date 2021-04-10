package e1

import "fmt"

type Graduate struct {
	Student
}

func (p *Graduate) testing() {
	fmt.Println("大学生正在考试中...")
}