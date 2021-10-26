package worker

import "fmt"

type Worker struct {
	// 工人池
	WorkerPool chan chan Job
	// 任务队列
	JobChan chan Job
	// 退出信号
	QuitChan chan bool
}

func NewWorker(workerPool chan chan Job) *Worker {
	return &Worker{
		//WorkerPool: workerPool,
		JobChan:    make(chan Job),
		QuitChan:   make(chan bool),
	}
}

func (this *Worker) Start() {
	go func() {
		for {
			// 将当前 worker 注册到 worker 队列中。
			this.WorkerPool <- this.JobChan
			select {
			case job := <-this.JobChan:
				fmt.Println(job.Content)
			case <-this.QuitChan:
				return
			}
		}
	}()
}

func (this *Worker) Stop() {
	go func() {
		this.QuitChan <- true
	}()
}
