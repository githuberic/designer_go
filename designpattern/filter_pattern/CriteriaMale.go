package filter_pattern

type CriteriaMale struct {
}

func (this *CriteriaMale) MeetCriteria(persons []Person) []Person {
	var malePersons []Person
	for _, v := range persons {
		if v.Gender == "Male" {
			malePersons = append(malePersons, v)
		}
	}
	return malePersons
}
