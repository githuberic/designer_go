package e2

import (
	"log"
	"testing"
)

func TestDecorator(t *testing.T) {
	o := &OriCalculate{1}
	m := MutCalculate{o, 2}
	a := AddCalculate{o, 3}

	log.Printf("Oricalculate is %d \n MutCalculate is %d\n AddCalculate is %d \n", o.Cal(), m.Cal(), a.Cal())
}
