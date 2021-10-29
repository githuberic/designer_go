package e2

import "sync"

type SingleTon struct {
	name string
}

var singleTon *SingleTon
var once sync.Once

func GetInstance() *SingleTon {
	once.Do(func() {
		singleTon=&SingleTon{}
	})
	return singleTon
}
