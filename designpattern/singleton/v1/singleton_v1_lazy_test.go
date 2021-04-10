package v1

import (
	"sync"
	"testing"
)
var wg sync.WaitGroup

func TestVerifyLazy(t *testing.T) {
	singletonList := make([]*SingleObject, 0)

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(group sync.WaitGroup) {
			defer wg.Done()
			singletonList = append(singletonList, GetInstanceV3())
		}(wg)
	}
	wg.Wait()

	for i := 0; i < len(singletonList)-1; i++ {
		if singletonList[i] != singletonList[i+1] {
			t.Fatal("instance is not equal")
		}
	}
}