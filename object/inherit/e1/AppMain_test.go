package e1

import "testing"

func TestVerify(t *testing.T)  {
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
