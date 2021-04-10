package good

/**
将数据库连接抽象为接口, 以支持多种数据库
*/
type IDBConnection interface {
	Execute(sql string, args ...interface{}) (error, int)
}
