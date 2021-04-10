package v1

import (
	"sync"
)

/**
双锁
*/
var lock *sync.Mutex = &sync.Mutex{}

func GetInstanceV3() *SingleObject {
	lock.Lock()
	defer lock.Unlock()
	if singleObj == nil {
		singleObj = new(SingleObject)
	}
	return singleObj
}

/**
双锁进阶 还可使用原子load以及赋值
*/
func GetInstanceV4() *SingleObject {
	if singleObj == nil {
		lock.Lock()
		defer lock.Unlock()
		singleObj = new(SingleObject)
	}
	return singleObj
}
