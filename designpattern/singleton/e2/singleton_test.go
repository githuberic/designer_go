package e2

import (
	"log"
	"sync"
	"testing"
)

func TestVerify(t *testing.T) {
	s1 := GetInstance()
	s2 := GetInstance()

	if s1 == s2 {
		log.Fatalf("s1 的地址是 %p,s2 的地址是 %p", s1, s2)
		log.Println("s1 和 s2 是相同的实例")
	} else {
		log.Println("s1 和 s2 是不相同的实例")
	}
}

var wg sync.WaitGroup

func TestVerifyV2(t *testing.T) {
	singletonList := make([]*SingleTon, 0, 100)

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(group sync.WaitGroup) {
			defer wg.Done()
			singletonList = append(singletonList, GetInstance())
		}(wg)
	}
	wg.Wait()

	for i := 0; i < len(singletonList)-1; i++ {
		if singletonList[i] != singletonList[i+1] {
			t.Fatal("instance is not equal")
		}
	}
}
