package bad

import (
	"testing"
)

func TestVerify(t *testing.T) {
	bl := NewBadTeamLeader(1, "张Leader")
	bl.CountOpeningTasks()
}
