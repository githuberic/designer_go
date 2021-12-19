package filter_pattern

type CriteriaSingle struct {
}

//按照未婚过滤
func (s *CriteriaSingle) MeetCriteria(persons []Person) []Person {
	var femalePersons []Person
	for _, person := range persons {
		if person.MaritalStatus == "Single" {
			femalePersons = append(femalePersons, person)
		}
	}
	return femalePersons
}
