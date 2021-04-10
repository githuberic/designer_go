package good

import "testing"

func TestVerify(t *testing.T)  {
	gu := NewGoodUser(2, "lgq")
	gu.Study(NewGolangCourse())

	gn := NewGoodUser(3, "eric")
	gn.Study(NewJavaCourse())
}
