package v2

type AliyunClientAdapter struct {
	Client AliyunClient
}

func (p *AliyunClientAdapter) CreateServer(cpu, mem float64) error {
	p.Client.CreateServer(int(cpu), int(mem))
	return nil
}
