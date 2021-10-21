package bad

import (
	"fmt"
	"designer_go/principle/law_of_demeter/common"
)

type TeamLeader struct {
	iID int
	sName string
}

func NewBadTeamLeader(id int, name string) *TeamLeader {
	return &TeamLeader{
		id,
		name,
	}
}

// 不好的ITeamLeader实现, 同时耦合了Task和BadTeamMember两个类
func (p *TeamLeader) CountOpeningTasks() int {
	tasks := common.LoadTaskList()
	member := NewBadTeamMember(11, "王Member")
	sum := member.countOpeningTasks(tasks)

	fmt.Printf("%v CountOpeningTasks, got %v", p.sName, sum)
	return sum
}

