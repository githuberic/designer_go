package good

import "testing"

func TestGood(t *testing.T) {
	fnCallAndLog := func(fn func() error) {
		err := fn()
		if err != nil {
			t.Logf("error = %s", err.Error())
		}
	}

	fnTestGoodBird := func(gb IGoodBird) {
		fnCallAndLog(gb.Tweet)
		if it, ok := gb.(IFlyableBird);ok {
			fnCallAndLog(it.Fly)
		}
		if it, ok := gb.(IRunnableBird);ok {
			fnCallAndLog(it.Run)
		}
	}

	fnTestGoodBird(NewGoodFlyableBird(11, "飞鸟"))
	fnTestGoodBird(NewGoodOstrichBird(12, "鸵鸟"))
}

