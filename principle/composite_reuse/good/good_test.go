package good

import (
	"testing"
	"designer_go/principle/composite_reuse/entity"
)

func TestVerifyGood(t *testing.T) {
	fnCallAndLog := func(fn func() (error, int)) {
		e, rows := fn()
		if e != nil {
			t.Errorf("error = %s", e.Error())
		} else {
			t.Logf("affected rows = %v", rows)
		}
	}

	con := NewMysqlConnection("database connection url", "sa", "123")
	gd := NewProductDAO()
	gd.SetDBConnection(con)

	product := entity.NewProduct(1, "手机", 1000)
	fnCallAndLog(func() (error, int) {
		return gd.Insert(product)
	})
	fnCallAndLog(func() (error, int) {
		return gd.Update(product)
	})
	fnCallAndLog(func() (error, int) {
		return gd.Delete(product.ID)
	})
}

