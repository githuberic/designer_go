package e1

import (
	"fmt"
)

type StopState struct{}

// NewStartState 实例化开始状态类
func NewStopState() *StopState {
	return &StopState{}
}

// DoAction 开始状态类的DoAction，实现State接口
func (start *StopState) DoAction(context *Context) {
	fmt.Println("Now is stop state")
	context.state = start
}

//ToString 返回开始状态类名称
func (start *StopState) ToString() string {
	return "Stop state"
}
