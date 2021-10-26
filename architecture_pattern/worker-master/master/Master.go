package master

import (
	"designer_go/architecture_pattern/worker-master/job"
	"designer_go/architecture_pattern/worker-master/worker"
)

type Master struct {
	WorkerPool chan chan job.Job   //worker池
	Result     map[interface{}]int //存放worker处理后的结果集
	jobQueue   chan job.Job        //待处理的任务chan
	workList   []worker.Worker     //存放worker列表，用于停止worker
}

//maxWorkers:开启线程数
var maxworker int

func NewMaster(maxWorkers int, result map[interface{}]int) *Master {
	pool := make(chan chan job.Job, maxWorkers)
	maxworker = maxWorkers
	return &Master{WorkerPool: pool,
		Result:   result,
		jobQueue: make(chan job.Job, 2*maxWorkers),
	}
}

func (m *Master) Run() {
	for i := 0; i < maxworker; i++ {
		work := worker.NewWorker(m.WorkerPool, m.Result, i)
		m.workList = append(m.workList, work)
		work.Start()
	}
	go m.dispatch()
}

func (m *Master) dispatch() {
	for {
		select {
		case jobQ := <-m.jobQueue:
			go func(jb job.Job) {
				// 从workerPool中取出一个worker的JobChannel
				jobChannel := <-m.WorkerPool
				// 向这个JobChannel中发送job，worker中的接收配对操作会被唤醒
				jobChannel <- jb
			}(jobQ)
		}
	}
}

// 添加任务到任务通道
func (m *Master) AddJob(num int) {
	job := job.NewJob(num)
	// 向任务通道发送任务
	m.jobQueue <- job
}

// 停止所有任务
func (m *Master) Stop() {
	for _, v := range m.workList {
		v.Start()
	}
}
