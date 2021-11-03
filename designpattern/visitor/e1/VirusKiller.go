package e1

import "fmt"

type VirusKiller struct{}

func (this *VirusKiller) VisitApartment(a Apartment) {
	fmt.Printf("Kill virus for Apartment %s\n", a.Name)
}
func (this *VirusKiller) VisitBuilding(b Building) {
	for e := b.Apartments.Front(); e != nil; e = e.Next() {
		this.VisitApartment(e.Value.(Apartment))
	}
}
