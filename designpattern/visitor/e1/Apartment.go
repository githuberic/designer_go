package e1

type Apartment struct {
	Name string
}
func (p Apartment) Accept(v Visitor) {
	v.VisitApartment(p)
}

