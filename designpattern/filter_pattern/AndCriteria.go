package filter_pattern

type AndCriteria struct {
	criteria      Criteria
	otherCriteria Criteria
}

func (this *AndCriteria) AndCriteria(criteria Criteria, otherCriteria Criteria) {
	this.criteria = criteria
	this.otherCriteria = otherCriteria
}

func (this *AndCriteria) MeetCretiria(persons []Person) []Person {
	group1 := this.criteria.MeetCriteria(persons)
	return this.otherCriteria.MeetCriteria(group1)
}
