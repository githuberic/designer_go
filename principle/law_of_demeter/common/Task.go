package common

type TaskStatus int

/**
任务状态
*/
const (
	OPENING = iota
	DONE
	CANCLED
	DENIED
)

// 定义任务信息, 以及加载任务清单的方法
type Task struct {
	iID     int
	iStatus TaskStatus
}

func NewTask(id int, status TaskStatus) *Task {
	return &Task{
		id,
		status,
	}
}

func (p *Task) ID() int {
	return p.iID
}
func (p *Task) Status() TaskStatus {
	return p.iStatus
}

func LoadTaskList() []*Task {
	tasks := make([]*Task, 0)
	tasks = append(tasks, NewTask(1, OPENING))
	tasks = append(tasks, NewTask(2, DONE))
	tasks = append(tasks, NewTask(3, CANCLED))
	tasks = append(tasks, NewTask(4, DENIED))
	return tasks
}
