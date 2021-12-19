package e3

import "fmt"

type IpNetWork struct {
	server *Server
}

func (p *IpNetWork) Conn(ip string) {
	p.server = &Server{localIP: ip}
	fmt.Println(p.server.localIP + " Connected")
}
func (p *IpNetWork) Close() {
	fmt.Println(p.server.localIP + " Closed")
}
