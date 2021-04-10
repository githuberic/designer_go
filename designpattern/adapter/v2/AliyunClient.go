package v2

import "fmt"

type AliyunClient struct {
}

func (p *AliyunClient) CreateServer(cpu, mem int) error {
	fmt.Printf("aliyun client run success, cpu： %d, mem: %d\n", cpu, mem)
	return nil
}
