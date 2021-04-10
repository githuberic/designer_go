package slice_test

import (
	"bytes"
	"fmt"
	"testing"
)

func TestSlice1(t *testing.T) {
	foo := make([]int, 5)
	foo[3] = 42
	foo[4] = 100

	bar := foo[1:4]
	bar[1] = 99
	// foo 和 bar 的内存是共享的，所以，foo 和 bar 对数组内容的修改都会影响到对方
	fmt.Println(foo)
	fmt.Println(bar)

	arrA := make([]int, 32)
	arrB := arrA[1:16]
	// 对 a 做了一个 append()的操作，这个操作会让 a 重新分配内存，这就会导致 a 和 b 不再共享
	arrA = append(arrA, 1)
	arrA[3] = 42
	fmt.Println(arrA)
	fmt.Println(arrB)
}

func TestSlice2(t *testing.T)  {
	path := []byte("AAAA/BBBBBBBBB")
	sepIndex := bytes.IndexByte(path,'/')

	//dir1 := path[:sepIndex:sepIndex]
	dir1 := path[:sepIndex]
	dir2 := path[sepIndex+1:]

	fmt.Println("dir1 =>",string(dir1)) //prints: dir1 => AAAA
	fmt.Println("dir2 =>",string(dir2)) //prints: dir2 => BBBBBBBBB

	dir1 = append(dir1,"suffix"...)
	fmt.Println("dir1 =>",string(dir1)) //prints: dir1 => AAAAsuffix
	fmt.Println("dir2 =>",string(dir2)) //prints: dir2 => uffixBBBB
}
