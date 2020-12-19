package memo_test

import (
	"testing"
	"tortoise/gopl.io/ch9/memotest"
	"tortoise/gopl.io/ch9/memo5"
)

var httpGetBody = memotest.HTTPGetBody

func Test(t *testing.T) {
	m := memo.New(httpGetBody)
	defer m.Close()
	memotest.Sequential(t, m)
}

func TestConcurrent(t *testing.T) {
	m := memo.New(httpGetBody)
	defer m.Close()
	memotest.Concurrent(t, m)
}