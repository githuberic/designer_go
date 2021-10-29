package e3

type NetWork interface {
	Conn(ip string)
	Close()
}
