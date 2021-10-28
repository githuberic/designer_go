package filter_pattern

type CriteriaFemale struct {
}

func (this *CriteriaFemale) MeetCriteria(persons []Person) []Person {
	var femalePersons []Person
	for _, v := range persons {
		if v.Gender == "Female" {
			femalePersons = append(femalePersons, v)
		}
	}
	return femalePersons
}
