package good

import "fmt"

type TeamLeader struct {
	iID int
	sName string
}

func NewGoodTeamLeader(id int, name string) *TeamLeader {
	return &TeamLeader{
		id,
		name,
	}
}

func (p *TeamLeader) CountOpeningTasks() int {
	member := NewGoodTeamMember(11, "王Member")
	sum := member.countOpeningTasks()

	fmt.Printf("%v CountOpeningTasks, got %v", p.sName, sum)
	return sum
}
