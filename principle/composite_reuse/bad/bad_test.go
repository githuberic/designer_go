package bad

import (
	"testing"
	"designer_go/principle/composite_reuse/entity"
)

func TestVerifyBad(t *testing.T) {
	fnCallAndLog := func(fn func() (error, int)) {
		e, rows := fn()
		if e != nil {
			t.Errorf("error = %s", e.Error())
		} else {
			t.Logf("affected rows = %v", rows)
		}
	}

	product := entity.NewProduct(1, "手机", 1000)
	bd := NewProductDAO("database connection url", "sa", "123")
	fnCallAndLog(func() (error, int) {
		return bd.Insert(product)
	})
	fnCallAndLog(func() (error, int) {
		return bd.Update(product)
	})
	fnCallAndLog(func() (error, int) {
		return bd.Delete(product.ID)
	})
}

