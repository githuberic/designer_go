package e2

type AwsClientAdapter struct {
	Client AWSClient
}

func (p *AwsClientAdapter) CreateServer(cpu, mem float64) error {
	p.Client.CreateServer(cpu, mem)
	return nil
}
