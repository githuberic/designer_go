package e2

type Week interface {
	Today()
	Next(*DayContext)
}
