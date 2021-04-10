package v1

import (
	"log"
	"sync"
	"testing"
)

func TestVerifyHungryV1(t *testing.T) {
	s1 := GetInstanceV1()
	s2 := GetInstanceV1()

	if s1 == s2 {
		log.Fatalf("s1 的地址是 %p,s2 的地址是 %p", s1, s2)
		log.Println("s1 和 s2 是相同的实例")
	} else {
		log.Println("s1 和 s2 是不相同的实例")
	}
}

func TestVerifyHungryV2(t *testing.T) {
	s1 := GetInstanceV2()
	s2 := GetInstanceV2()

	if s1 == s2 {
		log.Fatalf("s1 的地址是 %p,s2 的地址是 %p", s1, s2)
		log.Println("s1 和 s2 是相同的实例")
	} else {
		log.Println("s1 和 s2 是不相同的实例")
	}
}


func TestVerifyHungry(t *testing.T) {
	singletonList := make([]*SingleObject, 0)

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(group sync.WaitGroup) {
			defer wg.Done()
			singletonList = append(singletonList, GetInstanceV1())
		}(wg)
	}
	wg.Wait()

	for i := 0; i < len(singletonList)-1; i++ {
		if singletonList[i] != singletonList[i+1] {
			t.Fatal("instance is not equal")
		}
	}
}