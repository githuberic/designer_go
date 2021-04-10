package visitor

import (
	"container/list"
	"fmt"
	"testing"
)

type Target interface {
	Accept(Visitor)
}

type Visitor interface {
	VisitBuilding(building Building)
	VisitApartment(apartment Apartment)
}


type Apartment struct {
	Name string
}
func (p Apartment) Accept(v Visitor) {
	v.VisitApartment(p)
}

type Building struct {
	Apartments list.List
}
func (b Building) Accept(v Visitor) {
	v.VisitBuilding(b)
}


type Cleaner struct {}

func (p *Cleaner) VisitApartment(a Apartment)  {
	fmt.Printf("cleanup Apartment %s\n", a.Name)
}
func (p *Cleaner) VisitBuilding(b Building) {
	for e := b.Apartments.Front(); e != nil; e = e.Next() {
		p.VisitApartment(e.Value.(Apartment))
	}
}

type VirusKiller struct {}
func (vk VirusKiller) VisitApartment(a Apartment) {
	fmt.Printf("Kill virus for Apartment %s\n", a.Name)
}
func (vk VirusKiller) VisitBuilding(b Building) {
	for e := b.Apartments.Front(); e != nil; e = e.Next() {
		vk.VisitApartment(e.Value.(Apartment))
	}
}

func TestVerify(t *testing.T)  {
	list:= list.List{}
	list.PushBack(Apartment{Name:"B101"})
	list.PushBack(Apartment{Name:"B102"})
	list.PushBack(Apartment{Name:"B103"})

	targets := []Target{
		Apartment{Name:"101"},
		Apartment{Name:"103"},
		Building{Apartments: list},
		Apartment{Name:"201"},
	}

	virusKiller := &VirusKiller{}//new(VirusKiller)
	for _, item := range targets {
		item.Accept(virusKiller)
	}


	cleaner := &Cleaner{}//new(Cleaner)
	for _, item := range targets {
		item.Accept(cleaner)
	}
}
// https://studygolang.com/articles/26513