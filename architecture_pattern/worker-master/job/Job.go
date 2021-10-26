package job

type Job struct {
	Num int
}

func NewJob(num int) Job {
	return Job{Num: num}
}
