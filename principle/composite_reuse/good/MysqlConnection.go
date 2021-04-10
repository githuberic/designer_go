package good

import "fmt"

/**
更好的设计, GoodMysqlConnection封装MYSQL数据库方言, 实现IGoodDBConnection接口
 */
type MysqlConnection struct {
	sURL string
	sUID string
	sPWD string
}

func NewMysqlConnection(uri string, uid string, pwd string) *MysqlConnection {
	return &MysqlConnection{
		sURL: uri,
		sUID: uid,
		sPWD: pwd,
	}
}

func (p *MysqlConnection) Execute(sql string, args ...interface{}) (error, int)  {
	fmt.Printf("GoodMysqlConnection.Execute, sql=%v, args=%v\n", sql, args)
	return nil, 0
}
