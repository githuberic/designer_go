package e1

import "fmt"

//编写一个学生考试系统
type Student struct {
	Name string
	Age int
	Score int
}

// 将Pupil和Graduate共有的方法绑定到 *Student
func (stu *Student) ShowInfo() {
	fmt.Printf("Name：%v Age：%v Score：%v \n",stu.Name,stu.Age,stu.Score)
}

func (stu *Student) SetScore (score int){
	stu.Score = score
}