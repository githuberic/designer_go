package filter_pattern

type OrCriteria struct {
	criteria      Criteria
	otherCriteria Criteria
}

func (this *OrCriteria) OrCriteria(criteria Criteria, otherCriteria Criteria) {
	this.criteria = criteria
	this.otherCriteria = otherCriteria
}

func (this *OrCriteria) MeetCretiria(persons []Person) []Person {
	group1 := this.criteria.MeetCriteria(persons)
	group2 := this.otherCriteria.MeetCriteria(persons)

	groupM := make([]Person, 0, len(group1)+len(group2))
	groupM = append(groupM, group1...)
	groupM = append(groupM, group2...)
	return groupM
}

func ArrayMerge(arrs ...[]interface{}) []interface{} {
	n := 0
	for _, v := range arrs {
		n += len(v)
	}

	arrN := make([]interface{}, 0, n)
	for _, v := range arrs {
		arrN = append(arrN, v...)
	}
	return arrN
}
