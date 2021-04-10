package good

import "fmt"

type TXYunConnection struct {
	sURL  string
	sUID  string
	sPWD  string
	token string
}

func NewTXYunConnection(uri string, uid string, pwd string) *TXYunConnection {
	return &TXYunConnection{
		sURL:  uri,
		sUID:  uid,
		sPWD:  pwd,
		token: "2312312312sdfadsf", //token生成
	}
}

func (p *TXYunConnection) Upload(path string, args ...interface{}) (error, int) {
	fmt.Printf("TXYunConnection.Upload, path=%v, args=%v\n", path, args)
	return nil, 0
}

func (p *TXYunConnection) Download(path string, args ...interface{}) (error, []byte) {
	fmt.Printf("TXYunConnection.Download, path=%v, args=%v\n", path, args)
	return nil, nil
}
