package compare_test

import (
	"fmt"
	"reflect"
	"testing"
)

// 深度比较
func TestCompare(t *testing.T) {
	m1 := map[string]string{"one": "a", "two": "b"}
	m2 := map[string]string{"two": "b", "one": "a"}
	fmt.Println("m1==m2",reflect.DeepEqual(m1,m2))

	s1 := []int{1,2,3}
	s2 := []int{1,2,3}
	fmt.Println("s1==s2",reflect.DeepEqual(s1,s2))
}
