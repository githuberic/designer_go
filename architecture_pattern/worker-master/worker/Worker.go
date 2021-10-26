package worker

import (
	"designer_go/architecture_pattern/worker-master/job"
	"fmt"
)

type Worker struct {
	id         int                 //workerId
	WorkerPool chan chan job.Job   //worker池
	JobChannel chan job.Job        //worker从JobChannel中获取Job进行处理
	Result     map[interface{}]int //worker将处理结果放入result
	quit       chan bool           //停止worker信号
}

func NewWorker(workerPool chan chan job.Job, result map[interface{}]int, id int) Worker {
	return Worker{
		id:         id,
		WorkerPool: workerPool,
		JobChannel: make(chan job.Job),
		Result:     result,
		quit:       make(chan bool),
	}
}

func (w Worker) Start() {
	go func() {
		for {
			//将worker的JobChannel放入master的workerPool中
			w.WorkerPool <- w.JobChannel
			select {
			//从JobChannel中获取Job进行处理，JobChannel是同步通道，会阻塞于此
			case job := <-w.JobChannel:
				x := job.Num * job.Num
				fmt.Println(w.id, ":", x)
				w.Result[x] = w.id
			case <-w.quit:
				return
			}
		}
	}()
}

func (w Worker) Stop() {
	go func() {
		w.quit <- true
	}()
}
