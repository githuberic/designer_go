package function_option

import (
	"crypto/tls"
	"testing"
	"time"
)

type Config struct {
	Protocol string
	Timeout  time.Duration
	Maxconns uint
	TLS      *tls.Config
}
type Server struct {
	Addr string
	Port int
	Conf *Config
}

func NewServer(addr string, port int, conf *Config) (*Server, error) {
	return &Server{
		Addr: addr,
		Port: port,
		Conf: conf,
	}, nil
}

func TestFunctionOptionV0(t testing.T) {
	var conf = Config{Protocol: "tcp", Timeout: 60 * time.Duration(20), Maxconns: 200, TLS: nil}
	_, _ = NewServer("locahost", 9000, &conf)
}
