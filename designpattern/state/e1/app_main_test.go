package e1

import (
	"fmt"
	"testing"
)

func TestVerify(t *testing.T) {
	context := NewContext()

	startState := NewStartState()
	startState.DoAction(context)
	fmt.Println(context.GetState().ToString())

	stopState := NewStopState()
	stopState.DoAction(context)
	fmt.Println(context.GetState().ToString())
}
