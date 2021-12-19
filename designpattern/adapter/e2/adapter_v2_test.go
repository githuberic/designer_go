package e2

import "testing"

func TestVerify(t *testing.T) {
	var srv1 ICreateServer = &AwsClientAdapter{
		Client: AWSClient{},
	}
	srv1.CreateServer(1.0, 2.0)

	var srv2 ICreateServer = &AliyunClientAdapter{
		Client: AliyunClient{},
	}
	srv2.CreateServer(2.0, 4.0)
}
