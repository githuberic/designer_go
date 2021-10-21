package microkernel

import (
	"context"
)

type Collector interface {
	// 事件回传EventReceiver
	Init(evtReceiver EventReceiver) error
	// Collector 在不同的协程里面，启动时传递Context信息
	Start(agtCtx context.Context) error
	Stop() error
	Destroy() error
}
