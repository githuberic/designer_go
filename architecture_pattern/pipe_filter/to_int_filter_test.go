package pipefilter

import (
	"reflect"
	"testing"
)

func TestConvertInt(t *testing.T) {
	tif := NewToIntFilter()
	res, err := tif.Process([]string{"1", "2", "3"})
	if err != nil {
		t.Fatal(err)
	}

	excepted := []int{1, 2, 3}
	if !reflect.DeepEqual(excepted, res) {
		t.Fatalf("The expected is %v,the actual is %v", excepted, res)
	}
}

func TestWrongInputForTIF(t *testing.T) {
	tif := NewToIntFilter()
	_, err := tif.Process([]string{"1", "2.2", "3"})
	if err == nil {
		t.Fatal("An error is expected for wrong input")
	}
	_, err = tif.Process([]int{1, 2, 3})
	if err == nil {
		t.Fatal("An error is expected for wrong input")
	}
}
