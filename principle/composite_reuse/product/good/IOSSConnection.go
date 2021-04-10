package good

/**
将OSS连接抽象为接口, 以支持多个OSS云服务提供商
*/
type IOSSConnection interface {
	Upload(path string, args ...interface{}) (error, int)
	Download(path string, args ...interface{}) (error, []byte)
}
