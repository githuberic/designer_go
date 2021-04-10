package v2

import "fmt"

type AWSClient struct {
}

func (p *AWSClient) CreateServer(cpu, mem float64) error {
	fmt.Printf("aws client run success, cpu： %f, mem: %f\n", cpu, mem)
	return nil
}
