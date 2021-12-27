package e1

import (
	"testing"
)

//命令模式，客户端通过调用者，传递不同的命令，然后不同的接受者对此进行处理
func TestVerify(t *testing.T) {
	invoker := NewInvoker(NewShutDownTV())
	invoker.Call()

	invokerV0 := NewInvoker(NewOpenTV())
	invokerV0.Call()
}
