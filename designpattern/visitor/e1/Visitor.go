package e1

type Visitor interface {
	VisitBuilding(building Building)
	VisitApartment(apartment Apartment)
}
