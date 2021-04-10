package proxy

import (
	"fmt"
	"testing"
)

type NetWork interface {
	Conn(ip string)
	Close()
}

type Server struct {
	localIP string
}

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

type ProxyNetWork struct {
	ip *IpNetWork
}
func (p *ProxyNetWork) Conn(ip string)  {
	p.ip = &IpNetWork{}
	p.ip.Conn(ip)
}
func (p *ProxyNetWork) Close()  {
	p.ip.Close()
}

func TestVerify(t *testing.T)  {
	proxy := &ProxyNetWork{}
	proxy.Conn("192.168.10.2")
	proxy.Close()
}
