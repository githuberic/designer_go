package decoration

import (
	"fmt"
	"testing"
)

func decorator(fn func(s string)) func(s string) {
	return func(s string) {
		fmt.Println("Started")
		fn(s)
		fmt.Println("Done")
	}
}

func Hello(s string) {
	fmt.Println(s)
}

func TestVerifyV1(t *testing.T) {
	hello := decorator(Hello)
	hello("Hello")

	decorator(Hello)("Hello, World!")
}
