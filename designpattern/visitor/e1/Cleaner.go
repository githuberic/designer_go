package e1

import "fmt"

type Cleaner struct{}

func (this *Cleaner) VisitApartment(a Apartment) {
	fmt.Printf("Cleanup Apartment %s\n", a.Name)
}

func (this *Cleaner) VisitBuilding(b Building) {
	for e := b.Apartments.Front(); e != nil; e = e.Next() {
		this.VisitApartment(e.Value.(Apartment))
	}
}
