package good

import "fmt"

type QNYunConnection struct {
	sURL string
	sUID string
	sPWD string
}

func NewQNYunConnection(uri string, uid string, pwd string) *QNYunConnection {
	return &QNYunConnection{
		sURL: uri,
		sUID: uid,
		sPWD: pwd,
	}
}

func (p *QNYunConnection) Upload(path string, args ...interface{}) (error, int) {
	fmt.Printf("TXYunConnection.Upload, sql=%v, args=%v\n", path, args)
	return nil, 0
}

func (p *QNYunConnection) Download(path string, args ...interface{}) (error, []byte) {
	fmt.Printf("TXYunConnection.Download, sql=%v, args=%v\n", path, args)
	return nil, nil
}
