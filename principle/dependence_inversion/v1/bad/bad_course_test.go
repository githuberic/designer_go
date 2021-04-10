package bad

import "testing"

func TestVerify(t *testing.T)  {
	bu := NewBadUser(1, "Tom")
	bu.StudyGolangCourse()
}
