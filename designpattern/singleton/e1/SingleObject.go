package e1

import "sync"

type SingleObject struct {
	obj interface{}
}

var singleObj *SingleObject

/**
懒汉
*/
func GetInstanceV1() *SingleObject {
	if singleObj == nil {
		singleObj = new(SingleObject)
	}
	return singleObj
}

/**
饿汉
*/
func init() {
	singleObj = new(SingleObject)
}
func GetInstanceV2() *SingleObject {
	return singleObj
}

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
