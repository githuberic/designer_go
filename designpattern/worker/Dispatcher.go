package worker

// 调度器
type Dispatcher struct {
	WorkerPool chan chan Job
	MaxWorks   int
}

func NewDispatcher(maxWorks int) *Dispatcher {
	pool := make(chan chan Job, maxWorks)
	return &Dispatcher{WorkerPool: pool, MaxWorks: maxWorks}
}

func (this *Dispatcher) Run() {
	for i := 0; i < this.MaxWorks; i++ {
		worker := NewWorker(this.WorkerPool)
		worker.Start()
	}
	// 调度
	go this.dispatch()
}

var JobQueue = make(chan Job, 1000)

func (this *Dispatcher) dispatch() {
	for {
		select {
		case job := <-JobQueue:
			go func(job Job) {
				// 获取可用的worker通道，如果没有，则一直阻塞
				jobChanel := <-this.WorkerPool
				// 将job发送到worker任务通道
				jobChanel <- job
			}(job)
		}
	}
}
