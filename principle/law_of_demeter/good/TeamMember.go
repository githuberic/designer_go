package good

import (
	"designer_go/principle/law_of_demeter/common"
)

type TeamMember struct {
	iID int
	sName string
}

func NewGoodTeamMember(id int, name string) *TeamMember {
	return &TeamMember{
		id,
		name,
	}
}

func (p *TeamMember) countOpeningTasks() int {
	sum := 0
	tasks := common.LoadTaskList()

	for _,it := range tasks {
		if it.Status() == common.OPENING {
			sum++
		}
	}

	return sum
}