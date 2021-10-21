package good

import (
	"testing"
	"designer_go/principle/composite_reuse/product/entity"
)

func TestVerify(t *testing.T) {
	fnCallAndLog := func(fn func() (error, int)) {
		e, result := fn()
		if e != nil {
			t.Errorf("error = %s", e.Error())
		} else {
			t.Logf("result = %v", result)
		}
	}

	con := NewTXYunConnection("database connection url", "sa", "123")
	gd := NewProductImage()
	gd.SetOSSConnection(con)

	product := entity.NewProduct("翔龙武士宝剑", 899.00)
	fnCallAndLog(func() (error, int) {
		return gd.Upload(product)
	})
}
