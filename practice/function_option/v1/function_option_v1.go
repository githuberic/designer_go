package function_option

import (
	"crypto/tls"
	"time"
)

type Config struct {
	Protocol string
	Timeout  time.Duration
	Maxconns int
	TLS      *tls.Config
}
type Server struct {
	Addr string
	Port int
	Conf *Config
}


func NewServer(addr string, port int, conf *Config) (*Server, error) {
	//...
}

//Using the default configuratrion
srv1, _ := NewServer("localhost", 9000, nil)

conf := ServerConfig{Protocol:"tcp", Timeout: 60*time.Duration}
srv2, _ := NewServer("locahost", 9000, &conf)