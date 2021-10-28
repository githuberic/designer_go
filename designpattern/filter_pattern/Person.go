package filter_pattern

type Person struct {
	Name          string
	Gender        string
	MaritalStatus string
}

func GetPerson(name string, gender string, maritalStatus string) Person {
	return Person{Name: name, Gender: gender, MaritalStatus: maritalStatus}
}
