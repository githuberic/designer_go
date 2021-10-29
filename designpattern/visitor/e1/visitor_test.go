package e1

import (
	"container/list"
	"testing"
)

func TestVerify(t *testing.T) {
	list := list.List{}
	list.PushBack(Apartment{Name: "B101"})
	list.PushBack(Apartment{Name: "B102"})
	list.PushBack(Apartment{Name: "B103"})

	targets := []Target{
		Apartment{Name: "101"},
		Apartment{Name: "103"},
		Building{Apartments: list},
		Apartment{Name: "201"},
	}

	virusKiller := &VirusKiller{}
	for _, item := range targets {
		item.Accept(virusKiller)
	}

	cleaner := &Cleaner{}
	for _, item := range targets {
		item.Accept(cleaner)
	}
}
