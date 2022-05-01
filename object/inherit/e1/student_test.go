package e1

import (
	"fmt"
	"testing"
)

type Student struct {
	Name  string
	Age   int
	Score int
}

// 将Pupil和Graduate共有的方法绑定到 *Student
func (stu *Student) ShowInfo() {
	fmt.Printf("Name：%v Age：%v Score：%v \n", stu.Name, stu.Age, stu.Score)
}

func (stu *Student) SetScore(score int) {
	stu.Score = score
}

type Pupil struct {
	Student
}

func (p *Pupil) testing() {
	fmt.Println("小学生正在考试中...")
}

type Graduate struct {
	Student
}

func (p *Graduate) testing() {
	fmt.Println("大学生正在考试中...")
}

func TestVerify(t *testing.T) {
	pupil := &Pupil{}
	pupil.Student.Name = "tom"
	pupil.Student.Age = 8
	pupil.testing()
	pupil.SetScore(80)
	pupil.Student.ShowInfo()

	graduate := &Graduate{}
	graduate.Student.Name = "mary"
	graduate.Student.Age = 24
	graduate.testing()
	graduate.SetScore(95)
	graduate.Student.ShowInfo()
}
