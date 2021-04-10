package good

import "testing"

func TestVerify(t *testing.T)  {
	gl := NewGoodTeamLeader(2, "李Leader")
	gl.CountOpeningTasks()
}
