package bad

import (
	"designer_go/principle/law_of_demeter/common"
)

type TeamMember struct {
	iID   int
	sName string
}

func NewBadTeamMember(id int, name string) *TeamMember {
	return &TeamMember{
		id,
		name,
	}
}

func (p *TeamMember) countOpeningTasks(lstTasks []*common.Task) int {
	sum := 0
	for _, it := range lstTasks {
		if it.Status() == common.OPENING {
			sum++
		}
	}
	return sum
}
