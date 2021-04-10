package v1

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