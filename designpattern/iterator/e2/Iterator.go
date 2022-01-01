package e2

type Iterator interface {
	First()
	IsDone() bool
	Next() interface{}
}
