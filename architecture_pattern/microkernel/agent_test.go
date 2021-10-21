package microkernel

import (
	"fmt"
	"testing"
	"time"
)

func TestAgent(t *testing.T) {
	agt := NewAgent(100)
	c1 := NewCollect("c1", "1")
	c2 := NewCollect("c2", "2")

	agt.RegisterCollector("c1", c1)
	agt.RegisterCollector("c2", c2)

	if err := agt.Start(); err != nil {
		fmt.Printf("start error %v\n", err)
	}

	fmt.Println(agt.Start())
	time.Sleep(time.Second * 1)
	agt.Stop()
	agt.Destroy()
}
