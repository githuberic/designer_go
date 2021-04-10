package v2

import "testing"

/**
课程接口测试用例
*/
func TestVerify(t *testing.T) {
	fnShowCourse := func(it IProduct) {
		t.Logf("id=%v, name=%v, price=%v\n", it.ID(), it.Name(), it.Price())
	}

	c1 := NewProduct(1, "golang课程", 100)
	fnShowCourse(c1)

	c2 := NewDiscountProduct(2, "golang优惠课程", 100, 0.6)
	fnShowCourse(c2)

	c3 := NewDiscountProductV2(c1, 0.6)
	fnShowCourse(c3)
}
