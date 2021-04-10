package bad

import "testing"

func TestBad(t *testing.T) {
	fnCallAndLog := func(fn func() error) {
		err := fn()
		if err != nil {
			t.Logf("error = %s", err.Error())
		}
	}

	bb := NewBadNormalBird(1, "正常鸟(非鸵鸟)")
	fnCallAndLog(bb.Tweet)
	fnCallAndLog(bb.Fly)

	bo := NewBadOstrichBird(2, "鸵鸟")
	fnCallAndLog(bo.Tweet)
	fnCallAndLog(bo.Fly)
	if it, ok := bo.(*BadOstrichBird); ok {
		fnCallAndLog(it.Run)
	}
}
