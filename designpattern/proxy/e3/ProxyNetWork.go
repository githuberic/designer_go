package e3

type ProxyNetWork struct {
	ip *IpNetWork
}

func (p *ProxyNetWork) Conn(ip string) {
	p.ip = &IpNetWork{}
	p.ip.Conn(ip)
}
func (p *ProxyNetWork) Close() {
	p.ip.Close()
}
