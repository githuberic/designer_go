package worker_master

import (
	"designer_go/architecture_pattern/worker-master/master"
	"fmt"
	"testing"
)

func TestWorkerMaster(t *testing.T) {
	result := map[interface{}]int{}
	mas := master.NewMaster(5, result)
	mas.Run()

	for i := 0; i < 10; i++ {
		mas.AddJob(i)
	}

	//time.Sleep(time.Microsecond)
	fmt.Println("Result=", result)
}
