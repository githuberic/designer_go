package builder

import (
	"fmt"
	"testing"
)

func TestVerifyV1(t *testing.T) {
	director := new(Director)
	computer := director.Create("Intel","Kingsidun","WD")
	fmt.Println(*computer)
}
