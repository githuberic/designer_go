package e2

import "testing"

func TestProxy(t *testing.T) {
	sub := &Proxy{}
	res := sub.Do()

	if res != "pre:real:after" {
		t.Fail()
	}
}
