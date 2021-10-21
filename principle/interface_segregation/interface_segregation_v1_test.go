package interface_segregation

import (
	"testing"
	"designer_go/principle/interface_segregation/bad"
	"designer_go/principle/interface_segregation/good"
)

func TestVerifyBad(t *testing.T) {
	fnLogIfError := func(fn func() error) {
		e := fn()
		if e != nil {
			t.Logf("error = %s\n", e.Error())
		}
	}

	fnTestBadAnimal := func(a bad.IBadAnimal) {
		fnLogIfError(a.Eat)
		fnLogIfError(a.Fly)
		fnLogIfError(a.Swim)
	}

	fnTestBadAnimal(bad.NewBadBird(1, "BadBird"))
	fnTestBadAnimal(bad.NewBadDog(2, "BadDog"))
}

func TestVerifyGood(t *testing.T) {
	fnLogIfError := func(fn func() error) {
		e := fn()
		if e != nil {
			t.Logf("error = %s\n", e.Error())
		}
	}

	fnTestGoodAnimal := func(iGoodAnimal good.IGoodAnimal) {
		if it, ok := iGoodAnimal.(good.ISupportEat); ok {
			fnLogIfError(it.Eat)
		} else {
			t.Logf("%v/%v cannot eat", iGoodAnimal.Name(), iGoodAnimal.ID())
		}

		if it,ok := iGoodAnimal.(good.ISupportFly);ok {
			fnLogIfError(it.Fly)
		} else {
			t.Logf("%v/%v cannot fly", iGoodAnimal.Name(), iGoodAnimal.ID())
		}

		if it,ok := iGoodAnimal.(good.ISupportSwim);ok {
			fnLogIfError(it.Swim)
		} else {
			t.Logf("%v/%v cannot swim", iGoodAnimal.Name(), iGoodAnimal.ID())
		}
	}

	fnTestGoodAnimal(good.NewGoodBird(11, "GoodBird"))
	fnTestGoodAnimal(good.NewGoodDog(12, "GoodDog"))
}

// https://studygolang.com/articles/33103
