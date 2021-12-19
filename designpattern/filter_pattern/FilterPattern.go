package filter_pattern

type Criteria interface {
	MeetCriteria(persons []Person) []Person
}

