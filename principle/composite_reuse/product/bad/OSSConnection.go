package bad

import "fmt"

type OSSConnection struct {
	sURL string
	sUID string
	sPWD string
}

func NewOSSConnection(url string, uid string, pwd string) *OSSConnection {
	return &OSSConnection{
		url, uid, pwd,
	}
}

func (p *OSSConnection) Execute(sql string, args ...interface{}) (error, int) {
	fmt.Printf("OSSConnection.Execute, sql=%v, args=%v\n", sql, args)
	return nil, 0
}
