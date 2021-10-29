package e1

import "fmt"

type VirusKiller struct{}

func (vk VirusKiller) VisitApartment(a Apartment) {
	fmt.Printf("Kill virus for Apartment %s\n", a.Name)
}
func (vk VirusKiller) VisitBuilding(b Building) {
	for e := b.Apartments.Front(); e != nil; e = e.Next() {
		vk.VisitApartment(e.Value.(Apartment))
	}
}
