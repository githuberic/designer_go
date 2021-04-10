package bad

import "fmt"

// BadDBConnection用于连接数据库并执行SQL语句
type DBConnection struct {
	sURL string
	sUID string
	sPWD string
}

func NewDBConnection(url string, uid string, pwd string) *DBConnection {
	return &DBConnection{
		url, uid, pwd,
	}
}
func (p *DBConnection) Execute(sql string, args... interface{}) (error, int) {
	fmt.Printf("BadDBConnection.Execute, sql=%v, args=%v\n", sql, args)
	return nil,0
}